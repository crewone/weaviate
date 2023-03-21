//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package acceptance_with_go_client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	client "github.com/weaviate/weaviate-go-client/v4/weaviate"
)

func TestAutoschemaCasingClass(t *testing.T) {
	ctx := context.Background()
	c := client.New(client.Config{Scheme: "http", Host: "localhost:8080"})

	upperClassName := "RandomTestClass1234"
	lowerClassName := "randomTestClass1234"

	cases := []struct {
		className1 string
		className2 string
	}{
		{className1: upperClassName, className2: upperClassName},
		{className1: lowerClassName, className2: lowerClassName},
		{className1: upperClassName, className2: lowerClassName},
		{className1: lowerClassName, className2: upperClassName},
	}
	for _, tt := range cases {
		t.Run(tt.className1+" "+tt.className2, func(t *testing.T) {
			c.Schema().ClassDeleter().WithClassName(tt.className1).Do(ctx)
			c.Schema().ClassDeleter().WithClassName(tt.className2).Do(ctx)
			creator := c.Data().Creator()
			_, err := creator.WithClassName(tt.className1).Do(ctx)
			require.Nil(t, err)

			_, err = creator.WithClassName(tt.className2).Do(ctx)
			require.Nil(t, err)

			// class exists only once in Uppercase, so lowercase delete has to fail
			require.Nil(t, c.Schema().ClassDeleter().WithClassName(upperClassName).Do(ctx))
			require.NotNil(t, c.Schema().ClassDeleter().WithClassName(lowerClassName).Do(ctx))
		})
	}
}
