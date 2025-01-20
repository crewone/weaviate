//go:build ignore
// +build ignore

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/oklog/run"
	"github.com/sirupsen/logrus"
	"github.com/weaviate/weaviate/adapters/repos/db/lsmkv"
	"github.com/weaviate/weaviate/adapters/repos/db/lsmkv/segmentindex"
	"github.com/weaviate/weaviate/entities/cyclemanager"
	"gopkg.in/yaml.v3"
)

func main() {
	// logger
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{})

	// flags
	var (
		target            string
		runtimeConfig     string
		dataPath          string
		backupDestination string
	)
	flag.StringVar(&target, "target", "all", "target running mode (all, api-server, backup-server, retention-server")
	flag.StringVar(&runtimeConfig, "runtime-config", "runtime-overrides.yaml", "runtime config that will be reloaded every second")
	flag.StringVar(&dataPath, "path", "data", "data path for segment files")
	flag.StringVar(&backupDestination, "backup-dst", "backup", "destination to copy the backups to")
	flag.Parse()

	// run group lifecycle
	var g run.Group
	ctx, cancel := context.WithCancel(context.Background())

	cm := &RuntimeConfigManager{
		path:         runtimeConfig,
		reloadPeriod: 2 * time.Second,
	}

	{
		// manage runtime confimanager
		g.Add(func() error {
			cm.loop()
			return nil
		}, func(err error) {
			logger.Info("runtime configmanager stopped")
			cancel()
		})
	}

	if target == "api-server" || target == "all" {
		// api-server

		a := ApiServer{
			logger:   logger,
			dataPath: dataPath,
			bc:       lsmkv.NewBucketCreator(),
			cm:       cm,
		}
		g.Add(func() error {
			return a.Run(ctx)
		}, func(err error) {
			logger.Info("api server stopped")
			cancel()
		})
	}
	if target == "backup-server" || target == "all" {
		// backup-server
		a := BackupServer{
			logger: logger,
			cm:     cm,
		}

		// cron
		g.Add(func() error {
			return a.RunBackupCron(ctx)
		}, func(err error) {
			logger.Info("backup server stopped")
			cancel()
		})

		// http server
		// listener, err := net.Listen("tcp", "0.0.0.0:7070")
		// g.Add(func() error {
		// 	h := a.HTTPHandler(ctx)
		// 	if err != nil {
		// 		logger.Error("failed to create http listener")
		// 		return err
		// 	}
		// 	return http.Serve(listener, h)
		// }, func(err error) {
		// 	logger.Info("backup server stopped")
		// 	cancel()
		// 	listener.Close()
		// })

	}
	if target == "retention-server" || target == "all" {
		// retention-server
		a := RetentionServer{
			logger: logger,
			cm:     cm,
		}
		g.Add(func() error {
			return a.Run(ctx)
		}, func(err error) {
			logger.Info("retention server stopped")
			cancel()
		})
	}

	// Run all the actors
	if err := g.Run(); err != nil {
		logger.Info("stopping the server")
		os.Exit(1)
	}
}

// Think of core weaviate api server. The whole purpose of `ApiServer` here is to
// generate segment files without compaction (same assumption as Eitenie POC)
type ApiServer struct {
	logger   logrus.FieldLogger
	dataPath string
	bc       lsmkv.BucketCreator
	cm       *RuntimeConfigManager
}

func (a *ApiServer) Run(ctx context.Context) error {
	a.logger.Info("running api-server")
	flushCallbacks := cyclemanager.NewCallbackGroup("flush", a.logger, 1)
	flushCycle := cyclemanager.NewManager(cyclemanager.MemtableFlushCycleTicker(), flushCallbacks.CycleCallback, a.logger)
	flushCycle.Start()

	compactionCallbacks := cyclemanager.NewCallbackGroup("compactions", a.logger, 1)
	compactionCycle := cyclemanager.NewManager(cyclemanager.NewFixedTicker(4*time.Second), compactionCallbacks.CycleCallback, a.logger)
	compactionCycle.Start()

	h, m, s := time.Now().Clock()

	bucket, err := a.bc.NewBucket(ctx, filepath.Join(a.dataPath, "my-bucket"), "", a.logger, nil,
		compactionCallbacks, flushCallbacks,
		lsmkv.WithPread(true),
		lsmkv.WithForceCompation(true),
	)
	if err != nil {
		panic(err)
	}

	defer bucket.Shutdown(context.Background())

	// create phase segement files every 5s
	tk := time.NewTicker(10 * time.Second)
	phase := 0
	for {
		for i := 0; i < 100; i++ {
			key := []byte(fmt.Sprintf("phase-%02d-key-%03d", phase, i))
			value := []byte(fmt.Sprintf("written %02d:%02d:%02d", h, m, s))
			err := bucket.Put(key, value)
			if err != nil {
				panic(err)
			}
		}

		err, path := bucket.FlushAndSwitchX()
		if err != nil {
			panic(err)
		}
		f, err := os.Open(fmt.Sprintf("%s.db", path))
		if err != nil {
			panic(err)
		}
		defer f.Close()

		header, err := segmentindex.ParseHeader(f)
		if err != nil {
			panic(err)
		}

		if a.cm.getConfig().Compact.SkipMarker && header.Level == 0 {
			if _, err := os.Create(fmt.Sprintf("%s.%s", path, "skip-compact")); err != nil {
				panic(err)
			}
		}

		<-tk.C
		phase++
	}
}

// BackupServer handles copying segment files from `src` to `dst`
// This can be run as different container. But needs to share the pod because
// `src` is same filesystem where segments are written by `ApiServer` and need to be
// shared between `ApiServer` and `BackupServer`
type BackupServer struct {
	logger logrus.FieldLogger
	cm     *RuntimeConfigManager
}

// Responsible for triggering backup based on given backup policy
func (a *BackupServer) RunBackupCron(ctx context.Context) error {
	a.logger.Info("running backup cron server")

	t := time.NewTicker(2 * time.Second) // check for runtime config every 2 second
	for {
		<-t.C
		if !a.cm.getConfig().Backup.Enabled {
			a.logger.Info("backup is disabled. Skipping")
			continue
		}

		tt := time.NewTicker(a.cm.getConfig().Backup.Every)
		<-tt.C
		err := a.backup(ctx)
		if err != nil {
			a.logger.Error("backup failed", err)
			continue
		}
		a.logger.Info("backing up every", a.cm.getConfig().Backup.Every)
	}

	return nil
}

type BackupPolicy struct {
	Enabled bool          `yaml:"enabled"`
	Every   time.Duration `yaml:"every"`
}

// Responsible for user facing /v1/backup and /v1/restore.
func (a *BackupServer) HTTPHandler(ctx context.Context) http.Handler {
	a.logger.Info("running backup http server")
	return nil
}

// does the heavy job of backup.
func (a *BackupServer) backup(ctx context.Context) error {
	candidates := []string{}
	err := filepath.WalkDir("./data", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		// only copy the files that also has .skip-compact suffix
		if filepath.Ext(path) == ".db" {
			withoutExt := strings.TrimSuffix(path, filepath.Ext(".db"))
			s := fmt.Sprintf("%s.%s", withoutExt, "skip-compact")
			ok, err := fileExists(s)
			if err != nil {
				panic(err)
			}
			if ok {
				candidates = append(candidates, path)
			}
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	levelZeroSegments := []string{}
	for _, candidate := range candidates {
		f, err := os.Open(candidate)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		header, err := segmentindex.ParseHeader(f)
		if err != nil {
			panic(err)
		}

		if header.Level == 0 {
			levelZeroSegments = append(levelZeroSegments, candidate)
		}
	}

	segmentsCopied := 0
	filesCopied := 0
	for _, segment := range levelZeroSegments {
		withoutExt := strings.TrimSuffix(segment, filepath.Ext(segment))

		for _, ext := range []string{".db", ".bloom", ".cna"} {
			source := withoutExt + ext
			target := strings.Replace(withoutExt, "data/", "backup/", 1)
			target = target + ext

			os.MkdirAll(filepath.Dir(target), 0o755)
			cpCmd := exec.Command("cp", "-rf", source, target)
			cpCmd.Stdout = os.Stdout
			cpCmd.Stderr = os.Stderr
			err := cpCmd.Run()
			if err != nil {
				break
			}
			filesCopied++
		}
		segmentsCopied++
	}

	for _, segment := range candidates {
		withoutExt := strings.TrimSuffix(segment, filepath.Ext(segment))
		// segment copied. Now remove `skip-compact` marker file
		s := fmt.Sprintf("%s.%s", withoutExt, "skip-compact")
		rm := exec.Command("rm", "-rf", s)
		rm.Stdout = os.Stdout
		rm.Stderr = os.Stderr
		err := rm.Run()
		if err != nil {
			break
		}
	}

	fmt.Printf("Copied %d segments (%d files)\n", segmentsCopied, filesCopied)
	return nil
}

// does the heavy job of restore.
func (a *BackupServer) restore(ctx context.Context) error {
	a.logger.Info("restore done")
	return nil
}

// RetentionServer is takes care of segment files copied to `dst` by the backupserver.
// it's responsibilities are
// 1. How long the segment files lives?
// 2. When to merge these files to be able to restore backup like `hourly`, `weekly` and `monthly` retentions.
type RetentionServer struct {
	logger logrus.FieldLogger
	cm     *RuntimeConfigManager
}

type RetentionPolicy struct {
	Enabled bool
	Every   time.Duration
	Path    string
}

func (a *RetentionServer) Run(ctx context.Context) error {
	a.logger.Info("running retention-server")
	t := time.NewTicker(2 * time.Second) // check for runtime config every 2 second
	for {
		<-t.C
		if !a.cm.getConfig().Retention.Enabled {
			a.logger.Info("retention is disabled. Skipping")
			continue
		}

		tt := time.NewTicker(a.cm.getConfig().Retention.Every)
		<-tt.C
		err := a.compact(ctx)
		if err != nil {
			a.logger.Error("retention compact failed", err)
			continue
		}
		a.logger.Info("retention compacting every", a.cm.getConfig().Retention.Every)
	}
	return nil
}

// compacting the snapshot in time on retention server
func (a *RetentionServer) compact(ctx context.Context) error {
	cr := lsmkv.NewBucketCreator()

	dir := fmt.Sprintf("./%s", a.cm.getConfig().Retention.Path)

	flushCallbacks := cyclemanager.NewCallbackGroup("flush", a.logger, 1)
	flushCycle := cyclemanager.NewManager(cyclemanager.MemtableFlushCycleTicker(), flushCallbacks.CycleCallback, a.logger)
	flushCycle.Start()
	compactionCallbacks := cyclemanager.NewCallbackGroup("compactions", a.logger, 1)
	cyclemanager.NewManager(cyclemanager.NewFixedTicker(100*time.Millisecond), compactionCallbacks.CycleCallback, a.logger).Start()

	bucket, err := cr.NewBucket(ctx, filepath.Join(dir, "my-bucket"), "", a.logger, nil,
		compactionCallbacks, flushCallbacks,
		lsmkv.WithPread(true),
		lsmkv.WithForceCompation(true),
	)
	if err != nil {
		panic(err)
	}

	defer bucket.Shutdown(context.Background())

	time.Sleep(3 * time.Second)
	return nil
}

// Runtime config manager
type RuntimeConfigManager struct {
	reloadPeriod time.Duration
	path         string

	// mu protects the policy and keep it upto date with
	// runtime overrides
	mu     sync.Mutex // can be RWLock.
	config RuntimeConfig
}

func (cm *RuntimeConfigManager) loop() {
	tk := time.NewTicker(cm.reloadPeriod)

	for {
		f, err := os.Open(cm.path)
		if err != nil {
			panic(err)
		}
		b, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}
		var c RuntimeConfig
		if err := yaml.Unmarshal(b, &c); err != nil {
			panic(err)
		}

		cm.mu.Lock()
		cm.config = c
		cm.mu.Unlock()

		<-tk.C
	}
}

func (cm *RuntimeConfigManager) getConfig() *RuntimeConfig {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	return &cm.config
}

type RuntimeConfig struct {
	Compact   CompactPolicy   `yaml:"compact"`
	Backup    BackupPolicy    `yaml:"backup"`
	Retention RetentionPolicy `yaml:"retention"`
}

type CompactPolicy struct {
	SkipMarker bool `yaml:"skipmarker"`
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}

	return false, err
}
