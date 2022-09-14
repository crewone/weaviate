//go:build integrationTest
// +build integrationTest

package db

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/semi-technologies/weaviate/adapters/repos/db/vector/hnsw"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/usecases/config"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_DimensionTracking(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	dirName := t.TempDir()

	shardState := singleShardState()
	logger := logrus.New()
	schemaGetter := &fakeSchemaGetter{shardState: shardState}
	repo := New(logger, Config{
		RootPath:                  dirName,
		QueryMaximumResults:       10000,
		DiskUseWarningPercentage:  config.DefaultDiskUseWarningPercentage,
		DiskUseReadOnlyPercentage: config.DefaultDiskUseReadonlyPercentage,
		MaxImportGoroutinesFactor: 1,
	}, &fakeRemoteClient{},
		&fakeNodeResolver{}, nil)
	repo.SetSchemaGetter(schemaGetter)
	err := repo.WaitForStartup(testCtx())
	require.Nil(t, err)
	migrator := NewMigrator(repo, logger)

	t.Run("set schema", func(t *testing.T) {
		class := &models.Class{
			Class:               "Test",
			VectorIndexConfig:   hnsw.NewDefaultUserConfig(),
			InvertedIndexConfig: invertedConfig(),
		}
		schema := schema.Schema{
			Objects: &models.Schema{
				Classes: []*models.Class{class},
			},
		}

		require.Nil(t,
			migrator.AddClass(context.Background(), class, schemaGetter.shardState))

		schemaGetter.schema = schema
	})

	t.Run("import objects with d=128", func(t *testing.T) {
		dim := 128
		for i := 0; i < 100; i++ {
			vec := make([]float32, dim)
			for j := range vec {
				vec[j] = rand.Float32()
			}

			id := strfmt.UUID(uuid.MustParse(fmt.Sprintf("%032d", i)).String())
			obj := &models.Object{Class: "Test", ID: id}
			err := repo.PutObject(context.Background(), obj, vec)
			require.Nil(t, err)
		}
	})

	t.Run("import objects with d=0", func(t *testing.T) {
		for i := 100; i < 200; i++ {
			id := strfmt.UUID(uuid.MustParse(fmt.Sprintf("%032d", i)).String())
			obj := &models.Object{Class: "Test", ID: id}
			err := repo.PutObject(context.Background(), obj, nil)
			require.Nil(t, err)
		}
	})

	t.Run("verify dimensions after initial import", func(t *testing.T) {
		for _, shard := range repo.GetIndex("Test").Shards {
			assert.Equal(t, 12800, shard.Dimensions())
		}
	})

	t.Run("delete 10 objects with d=128", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			id := strfmt.UUID(uuid.MustParse(fmt.Sprintf("%032d", i)).String())
			err := repo.DeleteObject(context.Background(), "Test", id)
			require.Nil(t, err)
		}
	})

	t.Run("verify dimensions after delete", func(t *testing.T) {
		for _, shard := range repo.GetIndex("Test").Shards {
			assert.Equal(t, 11520, shard.Dimensions())
		}
	})

	t.Run("update some of the d=128 objects with a new vector", func(t *testing.T) {
		dim := 128
		for i := 0; i < 50; i++ {
			vec := make([]float32, dim)
			for j := range vec {
				vec[j] = rand.Float32()
			}

			id := strfmt.UUID(uuid.MustParse(fmt.Sprintf("%032d", i)).String())
			obj := &models.Object{Class: "Test", ID: id}
			// Put is idempotent, but since the IDs exist now, this is an update
			// under the hood and a "reinstert" for the already deleted ones
			err := repo.PutObject(context.Background(), obj, vec)
			require.Nil(t, err)
		}
	})

	t.Run("update some of the d=128 objects with a nil vector", func(t *testing.T) {
		for i := 50; i < 100; i++ {
			id := strfmt.UUID(uuid.MustParse(fmt.Sprintf("%032d", i)).String())
			obj := &models.Object{Class: "Test", ID: id}
			// Put is idempotent, but since the IDs exist now, this is an update
			// under the hood and a "reinstert" for the already deleted ones
			err := repo.PutObject(context.Background(), obj, nil)
			require.Nil(t, err)
		}
	})

	t.Run("verify dimensions after first set of updates", func(t *testing.T) {
		for _, shard := range repo.GetIndex("Test").Shards {
			// only half as many vectors as initially
			assert.Equal(t, 6400, shard.Dimensions())
		}
	})

	t.Run("update some of the origin nil vector objects with a d=128 vector", func(t *testing.T) {
		dim := 128
		for i := 100; i < 150; i++ {
			vec := make([]float32, dim)
			for j := range vec {
				vec[j] = rand.Float32()
			}

			id := strfmt.UUID(uuid.MustParse(fmt.Sprintf("%032d", i)).String())
			obj := &models.Object{Class: "Test", ID: id}
			// Put is idempotent, but since the IDs exist now, this is an update
			// under the hood and a "reinstert" for the already deleted ones
			err := repo.PutObject(context.Background(), obj, vec)
			require.Nil(t, err)
		}
	})

	t.Run("update some of the nil objects with another nil vector", func(t *testing.T) {
		for i := 150; i < 200; i++ {
			id := strfmt.UUID(uuid.MustParse(fmt.Sprintf("%032d", i)).String())
			obj := &models.Object{Class: "Test", ID: id}
			// Put is idempotent, but since the IDs exist now, this is an update
			// under the hood and a "reinstert" for the already deleted ones
			err := repo.PutObject(context.Background(), obj, nil)
			require.Nil(t, err)
		}
	})

	t.Run("verify dimensions after more updates", func(t *testing.T) {
		for _, shard := range repo.GetIndex("Test").Shards {
			// 000-050 d=128
			// 051-100 d=0
			// 100-150 d=128
			// 151-200 d=0
			assert.Equal(t, 12800, shard.Dimensions())
		}
	})

	// TODO WEAVIATE-286 still missing:
	//
	// - add objects without vectors using Batch Objects
	// - add objects with vectors using Batch Objects
	// - modify objects using Batch reference (this should not change the vector)
	// - modify object (attach a vector, remove a vector, replace a vector) using Merge
}
