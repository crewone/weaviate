//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
// 
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Package getmeta provides the local get meta graphql endpoint for Weaviate
package getmeta

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/semi-technologies/weaviate/adapters/handlers/graphql/descriptions"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/entities/schema/kind"
	"github.com/semi-technologies/weaviate/usecases/config"
)

// Build the local queries from the database schema.
func Build(dbSchema *schema.Schema, config config.Config) (*graphql.Field, error) {
	if len(dbSchema.Actions.Classes) == 0 && len(dbSchema.Things.Classes) == 0 {
		return nil, fmt.Errorf("there are no Actions or Things classes defined yet")
	}

	getMetaKinds := graphql.Fields{}
	if len(dbSchema.Actions.Classes) > 0 {
		localGetMetaActions, err := classFields(dbSchema.Actions.Classes, kind.Action, config)
		if err != nil {
			return nil, fmt.Errorf("failed to generate action fields from schema for local MetaGet because: %v", err)
		}

		getMetaKinds["Actions"] = &graphql.Field{
			Name:        "WeaviateLocalGetMetaActions",
			Description: descriptions.LocalGetMetaActions,
			Type:        localGetMetaActions,
			Resolve:     passThroughResolver,
		}
	}

	if len(dbSchema.Things.Classes) > 0 {
		localGetMetaThings, err := classFields(dbSchema.Things.Classes, kind.Thing, config)
		if err != nil {
			return nil, fmt.Errorf("failed to generate thing fields from schema for local MetaGet because: %v", err)
		}

		getMetaKinds["Things"] = &graphql.Field{
			Name:        "WeaviateLocalGetMetaThings",
			Description: descriptions.LocalGetMetaThings,
			Type:        localGetMetaThings,
			Resolve:     passThroughResolver,
		}
	}

	getMetaObj := graphql.NewObject(graphql.ObjectConfig{
		Name:        "WeaviateLocalGetMetaObj",
		Fields:      getMetaKinds,
		Description: descriptions.LocalGetObj,
	})

	localField := &graphql.Field{
		Name:        "WeaviateLocalGetMeta",
		Type:        getMetaObj,
		Description: descriptions.LocalGetMeta,
		Resolve:     passThroughResolver,
	}

	return localField, nil
}

func passThroughResolver(p graphql.ResolveParams) (interface{}, error) {
	// bubble up root resolver
	return p.Source, nil
}
