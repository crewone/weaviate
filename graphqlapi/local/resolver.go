/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */
package local

import (
	local_get "github.com/creativesoftwarefdn/weaviate/graphqlapi/local/get"
	local_getmeta "github.com/creativesoftwarefdn/weaviate/graphqlapi/local/getmeta"
)

// Resolving local GraphQL queries
type Resolver interface {
	local_get.Resolver
	local_getmeta.Resolver
}
