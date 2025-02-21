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

package mydist

// ShardDist shard distribution over nodes, copied from usecases/scaler.ShardDist to avoid import cycle
type ShardDist map[string][]string

// shards return names of all shards
func (m ShardDist) Shards() []string {
	ns := make([]string, 0, len(m))
	for node := range m {
		ns = append(ns, node)
	}
	return ns
}
