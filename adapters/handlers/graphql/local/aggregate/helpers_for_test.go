/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@semi.technology
 */
package aggregate

import (
	"context"
	"fmt"

	testhelper "github.com/semi-technologies/weaviate/adapters/handlers/graphql/test/helper"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/usecases/config"
	"github.com/semi-technologies/weaviate/usecases/kinds"
)

type mockRequestsLog struct{}

func (m *mockRequestsLog) Register(first string, second string) {

}

type mockResolver struct {
	testhelper.MockResolver
}

func newMockResolver(cfg config.Config) *mockResolver {
	field, err := Build(&testhelper.CarSchema, cfg)
	if err != nil {
		panic(fmt.Sprintf("could not build graphql test schema: %s", err))
	}
	mockLog := &mockRequestsLog{}
	mocker := &mockResolver{}
	mocker.RootFieldName = "Aggregate"
	mocker.RootField = field
	mocker.RootObject = map[string]interface{}{
		"Resolver":    Resolver(mocker),
		"RequestsLog": mockLog,
		"Config":      cfg,
	}

	return mocker
}

func (m *mockResolver) LocalAggregate(ctx context.Context, principal *models.Principal,
	params *kinds.AggregateParams) (interface{}, error) {
	args := m.Called(params)
	return args.Get(0), args.Error(1)
}
