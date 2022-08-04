//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package objects

import (
	"context"

	"github.com/pkg/errors"
	"github.com/semi-technologies/weaviate/adapters/repos/db/vector/hnsw"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/usecases/config"
	"github.com/sirupsen/logrus"
)

type vectorObtainer struct {
	vectorizerProvider VectorizerProvider
	schemaManager      schemaManager
	logger             logrus.FieldLogger
}

func newVectorObtainer(vectorizerProvider VectorizerProvider,
	schemaManager schemaManager, logger logrus.FieldLogger,
) *vectorObtainer {
	return &vectorObtainer{
		vectorizerProvider: vectorizerProvider,
		schemaManager:      schemaManager,
		logger:             logger,
	}
}

// Do retrieves the correct vector and makes sure it is set on the passed-in
// *models.Object. (This method mutates its paremeter)
func (vo *vectorObtainer) Do(ctx context.Context, obj *models.Object,
	principal *models.Principal,
) error {
	vectorizerName, cfg, err := vo.getVectorizerOfClass(obj.Class, principal)
	if err != nil {
		return err
	}

	hnswConfig, ok := cfg.(hnsw.UserConfig)
	if !ok {
		return errors.Errorf("vector index config (%T) is not of type HNSW, "+
			"but objects manager is restricted to HNSW", cfg)
	}

	if vectorizerName == config.VectorizerModuleNone {
		if err := vo.validateVectorPresent(obj, hnswConfig); err != nil {
			return NewErrInvalidUserInput("%v", err)
		}

		return nil
	}

	if hnswConfig.Skip {
		vo.logger.WithField("className", obj.Class).
			WithField("vectorizer", vectorizerName).
			Warningf("this class is configured to skip vector indexing, "+
				"but a vector was generated by the %q vectorizer. "+
				"This vector will be ignored. If you meant to index "+
				"the vector, make sure to set vectorIndexConfig.skip to 'false'. If the previous "+
				"setting is correct, make sure you set vectorizer to 'none' in the schema and "+
				"provide a null-vector (i.e. no vector) at import time.", vectorizerName)
	}

	// a vectorizer is present and indexing is not skipped
	// TODO: before calling the vectorizer we have to check if the user might have
	// overridden the vector
	if obj.Vector == nil {
		vectorizer, err := vo.vectorizerProvider.Vectorizer(vectorizerName, obj.Class)
		if err != nil {
			return err
		}

		if err := vectorizer.UpdateObject(ctx, obj); err != nil {
			return NewErrInternal("%v", err)
		}
	}

	return nil
}

func (vo *vectorObtainer) getVectorizerOfClass(className string,
	principal *models.Principal,
) (string, interface{}, error) {
	s, err := vo.schemaManager.GetSchema(principal)
	if err != nil {
		return "", nil, err
	}

	class := s.FindClassByName(schema.ClassName(className))
	if class == nil {
		// this should be impossible by the time this method gets called, but let's
		// be 100% certain
		return "", nil, errors.Errorf("class %s not present", className)
	}

	return class.Vectorizer, class.VectorIndexConfig, nil
}

func (vo *vectorObtainer) validateVectorPresent(obj *models.Object,
	hnswConfig hnsw.UserConfig,
) error {
	if hnswConfig.Skip && len(obj.Vector) > 0 {
		vo.logger.WithField("className", obj.Class).
			Warningf("this class is configured to skip vector indexing, " +
				"but a vector was explicitly provided. " +
				"This vector will be ignored. If you meant to index " +
				"the vector, make sure to set vectorIndexConfig.skip to 'false'. If the previous " +
				"setting is correct, make sure you set vectorizer to 'none' in the schema and " +
				"provide a null-vector (i.e. no vector) at import time.")
	}

	return nil
}
