//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package graphql

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"github.com/semi-technologies/weaviate/entities/modulecapabilities"
	"github.com/semi-technologies/weaviate/entities/schema/crossref"
	libvectorizer "github.com/semi-technologies/weaviate/usecases/vectorizer"
)

type Searcher struct {
	vectorizer libvectorizer.VectorizerDef
}

func NewSearcher(vectorizer libvectorizer.VectorizerDef) *Searcher {
	return &Searcher{vectorizer}
}

func (s *Searcher) VectorSearches() map[string]modulecapabilities.VectorForParams {
	vectorSearches := map[string]modulecapabilities.VectorForParams{}
	vectorSearches["nearText"] = s.vectorForNearTextParams
	return vectorSearches
}

func (s *Searcher) vectorForNearTextParams(ctx context.Context, params interface{},
	findVectorFn modulecapabilities.FindVectorFn) ([]float32, error) {
	return s.vectorFromNearTextParams(ctx,
		params.(*modulecapabilities.NearTextParams),
		findVectorFn,
	)
}

func (s *Searcher) vectorFromNearTextParams(ctx context.Context,
	params *modulecapabilities.NearTextParams, findVectorFn modulecapabilities.FindVectorFn) ([]float32, error) {
	vector, err := s.vectorizer.Corpi(ctx, params.Values)
	if err != nil {
		return nil, errors.Errorf("vectorize keywords: %v", err)
	}

	moveTo := params.MoveTo
	if moveTo.Force > 0 && (len(moveTo.Values) > 0 || len(moveTo.Objects) > 0) {
		moveToVector, err := s.vectorFromValuesAndObjects(ctx, moveTo.Values, moveTo.Objects, findVectorFn)
		if err != nil {
			return nil, errors.Errorf("vectorize move to: %v", err)
		}

		afterMoveTo, err := s.vectorizer.MoveTo(vector, moveToVector, moveTo.Force)
		if err != nil {
			return nil, err
		}
		vector = afterMoveTo
	}

	moveAway := params.MoveAwayFrom
	if moveAway.Force > 0 && (len(moveAway.Values) > 0 || len(moveAway.Objects) > 0) {
		moveAwayVector, err := s.vectorFromValuesAndObjects(ctx, moveAway.Values, moveAway.Objects, findVectorFn)
		if err != nil {
			return nil, errors.Errorf("vectorize move away from: %v", err)
		}

		afterMoveFrom, err := s.vectorizer.MoveAwayFrom(vector, moveAwayVector, moveAway.Force)
		if err != nil {
			return nil, err
		}
		vector = afterMoveFrom
	}

	return vector, nil
}

func (s *Searcher) vectorFromValuesAndObjects(ctx context.Context,
	values []string, objects []modulecapabilities.ObjectMove,
	findVectorFn modulecapabilities.FindVectorFn) ([]float32, error) {
	var objectVectors [][]float32

	if len(values) > 0 {
		moveToVector, err := s.vectorizer.Corpi(ctx, values)
		if err != nil {
			return nil, errors.Errorf("vectorize move to: %v", err)
		}
		objectVectors = append(objectVectors, moveToVector)
	}

	if len(objects) > 0 {
		var id strfmt.UUID
		for _, obj := range objects {
			if len(obj.ID) > 0 {
				id = strfmt.UUID(obj.ID)
			}
			if len(obj.Beacon) > 0 {
				ref, err := crossref.Parse(obj.Beacon)
				if err != nil {
					return nil, err
				}
				id = ref.TargetID
			}

			vector, err := findVectorFn(ctx, id)
			if err != nil {
				return nil, err
			}

			objectVectors = append(objectVectors, vector)
		}
	}

	return libvectorizer.CombineVectors(objectVectors), nil
}
