//go:build ignore
// +build ignore

package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/weaviate/weaviate/adapters/repos/db/lsmkv"
	"github.com/weaviate/weaviate/entities/cyclemanager"
)

func main() {
	ctx := context.Background()
	cr := lsmkv.NewBucketCreator()

	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	pathPtr := flag.String("path", "data", "specify path to inspect")
	flag.Parse()
	dir := fmt.Sprintf("./%s", *pathPtr)

	logger.Infof("inspecting data at %s", dir)

	flushCallbacks := cyclemanager.NewCallbackGroup("flush", logger, 1)
	flushCycle := cyclemanager.NewManager(cyclemanager.MemtableFlushCycleTicker(), flushCallbacks.CycleCallback, logger)
	flushCycle.Start()
	compactionCallbacks := cyclemanager.NewCallbackGroupNoop()

	bucket, err := cr.NewBucket(ctx, filepath.Join(dir, "my-bucket"), "", logger, nil,
		compactionCallbacks, flushCallbacks,
		lsmkv.WithPread(true),
	)
	if err != nil {
		panic(err)
	}

	defer bucket.Shutdown(context.Background())

	c := bucket.Cursor()
	defer c.Close()

	for k, v := c.First(); k != nil; k, v = c.Next() {
		fmt.Printf("key: %s, value: %s\n", k, v)
	}
}
