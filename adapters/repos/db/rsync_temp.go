//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package db

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/weaviate/weaviate/entities/backup"
	enterrors "github.com/weaviate/weaviate/entities/errors"
	"github.com/weaviate/weaviate/usecases/mydist"
)

// client the client interface is used to communicate with remote nodes
type client interface {
	PutFile(ctx context.Context, hostName, indexName,
		shardName, fileName string, payload io.ReadSeekCloser) error

	// CreateShard creates an empty shard on the remote node.
	// This is required in order to sync files to a specific shard on the remote node.
	CreateShard(ctx context.Context,
		hostName, indexName, shardName string) error

	// ReInitShard re-initialized new shard after all files has been synced to the remote node
	// Otherwise, it would not recognize the files when
	// serving traffic later.
	ReInitShard(ctx context.Context,
		hostName, indexName, shardName string) error
	IncreaseReplicationFactor(ctx context.Context, host, class string, dist mydist.ShardDist) error
}

// rsync synchronizes shards with remote nodes
type rsync struct {
	client          client
	cluster         nodeResolver
	persistenceRoot string
}

func newRSync(c client, cl nodeResolver, rootPath string) *rsync {
	return &rsync{client: c, cluster: cl, persistenceRoot: rootPath}
}

// Push pushes local shards of a class to remote nodes
func (r *rsync) Push(ctx context.Context, shardsBackups []*backup.ShardDescriptor, dist mydist.ShardDist, className string, logger logrus.FieldLogger) error {
	g := enterrors.NewErrorGroupWrapper(logger)
	g.SetLimit(_NUMCPU * 2)
	for _, desc := range shardsBackups {
		shardName := desc.Name
		additions := dist[shardName]
		desc := desc
		g.Go(func() error {
			return r.PushShard(ctx, className, desc, additions)
		}, shardName)

	}
	return g.Wait()
}

// PushShard replicates a shard on a set of nodes
func (r *rsync) PushShard(ctx context.Context, className string, desc *backup.ShardDescriptor, nodes []string) error {
	// Iterate over the new target nodes and copy files
	for _, node := range nodes {
		host, ok := r.cluster.NodeHostname(node)
		if !ok {
			return fmt.Errorf("%w: %q", "unresolved name", node)
		}
		if err := r.client.CreateShard(ctx, host, className, desc.Name); err != nil {
			return fmt.Errorf("create new shard on remote node %q: %w", node, err)
		}

		// simulate transferring large files (later have a sleep per file)
		// fmt.Println("NATEE simulating large files, sleeping for 3 seconds")
		// time.Sleep(3 * time.Second)
		// fmt.Println("NATEE done sleeping for simulation of large files")

		// Transfer each file that's part of the backup.
		for _, file := range desc.Files {
			err := r.PutFile(ctx, file, host, className, desc.Name)
			if err != nil {
				return fmt.Errorf("copy files to remote node %q: %w", node, err)
			}
		}

		// Transfer shard metadata files
		err := r.PutFile(ctx, desc.ShardVersionPath, host, className, desc.Name)
		if err != nil {
			return fmt.Errorf("copy shard version to remote node %q: %w", node, err)
		}

		err = r.PutFile(ctx, desc.DocIDCounterPath, host, className, desc.Name)
		if err != nil {
			return fmt.Errorf("copy index counter to remote node %q: %w", node, err)
		}

		err = r.PutFile(ctx, desc.PropLengthTrackerPath, host, className, desc.Name)
		if err != nil {
			return fmt.Errorf("copy prop length tracker to remote node %q: %w", node, err)
		}

		// Now that all files are on the remote node's new shard, the shard needs
		// to be reinitialized. Otherwise, it would not recognize the files when
		// serving traffic later.
		if err := r.client.ReInitShard(ctx, host, className, desc.Name); err != nil {
			return fmt.Errorf("create new shard on remote node %q: %w", node, err)
		}
	}
	return nil
}

func (r *rsync) PutFile(ctx context.Context, sourceFileName string,
	hostname, className, shardName string,
) error {
	absPath := filepath.Join(r.persistenceRoot, sourceFileName)
	f, err := os.Open(absPath)
	if err != nil {
		return fmt.Errorf("open file %q for reading: %w", absPath, err)
	}

	return r.client.PutFile(ctx, hostname, className, shardName, sourceFileName, f)
}
