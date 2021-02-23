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

package explore

import (
	"testing"

	"github.com/semi-technologies/weaviate/entities/search"
	"github.com/semi-technologies/weaviate/usecases/traverser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	name                      string
	query                     string
	expectedParamsToTraverser traverser.ExploreParams
	resolverReturn            []search.Result
	expectedResults           []result
}

type testCases []testCase

type result struct {
	pathToField   []string
	expectedValue interface{}
}

func Test_ResolveExplore(t *testing.T) {
	t.Parallel()

	testsNearText := testCases{
		testCase{
			name: "Resolve Explore with nearText",
			query: `
			{
					Explore(nearText: {concepts: ["car", "best brand"]}) {
							beacon className certainty
					}
			}`,
			expectedParamsToTraverser: traverser.ExploreParams{
				NearText: &traverser.NearTextParams{
					Values: []string{"car", "best brand"},
				},
			},
			resolverReturn: []search.Result{
				search.Result{
					Beacon:    "weaviate://localhost/some-uuid",
					ClassName: "bestClass",
					Certainty: 0.7,
				},
			},
			expectedResults: []result{{
				pathToField: []string{"Explore"},
				expectedValue: []interface{}{
					map[string]interface{}{
						"beacon":    "weaviate://localhost/some-uuid",
						"className": "bestClass",
						"certainty": float32(0.7),
					},
				},
			}},
		},

		testCase{
			name: "with nearText with optional limit and certainty set",
			query: `
			{
					Explore(
						nearText: {concepts: ["car", "best brand"], certainty: 0.6}, limit: 17 
						){
							beacon className
				}
			}`,
			expectedParamsToTraverser: traverser.ExploreParams{
				NearText: &traverser.NearTextParams{
					Values:    []string{"car", "best brand"},
					Certainty: 0.6,
				},
				Limit: 17,
			},
			resolverReturn: []search.Result{
				search.Result{
					Beacon:    "weaviate://localhost/some-uuid",
					ClassName: "bestClass",
				},
			},
			expectedResults: []result{{
				pathToField: []string{"Explore"},
				expectedValue: []interface{}{
					map[string]interface{}{
						"beacon":    "weaviate://localhost/some-uuid",
						"className": "bestClass",
					},
				},
			}},
		},

		testCase{
			name: "with moveTo set",
			query: `
			{
					Explore(
							limit: 17
							nearText: {
								concepts: ["car", "best brand"]
								moveTo: {
									concepts: ["mercedes"]
									force: 0.7
								}
							}
							) {
							beacon className
						}
			}`,
			expectedParamsToTraverser: traverser.ExploreParams{
				Limit: 17,
				NearText: &traverser.NearTextParams{
					Values: []string{"car", "best brand"},
					MoveTo: traverser.ExploreMove{
						Values: []string{"mercedes"},
						Force:  0.7,
					},
				},
			},
			resolverReturn: []search.Result{
				search.Result{
					Beacon:    "weaviate://localhost/some-uuid",
					ClassName: "bestClass",
				},
			},
			expectedResults: []result{{
				pathToField: []string{"Explore"},
				expectedValue: []interface{}{
					map[string]interface{}{
						"beacon":    "weaviate://localhost/some-uuid",
						"className": "bestClass",
					},
				},
			}},
		},

		testCase{
			name: "with moveTo and moveAwayFrom set",
			query: `
			{
					Explore(
							limit: 17
							nearText: {
								concepts: ["car", "best brand"]
								moveTo: {
									concepts: ["mercedes"]
									force: 0.7
								}
								moveAwayFrom: {
									concepts: ["van"]
									force: 0.7
								}
							}
							) {
							beacon className
						}
			}`,
			expectedParamsToTraverser: traverser.ExploreParams{
				Limit: 17,
				NearText: &traverser.NearTextParams{
					Values: []string{"car", "best brand"},
					MoveTo: traverser.ExploreMove{
						Values: []string{"mercedes"},
						Force:  0.7,
					},
					MoveAwayFrom: traverser.ExploreMove{
						Values: []string{"van"},
						Force:  0.7,
					},
				},
			},
			resolverReturn: []search.Result{
				search.Result{
					Beacon:    "weaviate://localhost/some-uuid",
					ClassName: "bestClass",
				},
			},
			expectedResults: []result{{
				pathToField: []string{"Explore"},
				expectedValue: []interface{}{
					map[string]interface{}{
						"beacon":    "weaviate://localhost/some-uuid",
						"className": "bestClass",
					},
				},
			}},
		},

		testCase{
			name: "with moveTo and objects set",
			query: `
			{
					Explore(
							limit: 17
							nearText: {
								concepts: ["car", "best brand"]
								moveTo: {
									concepts: ["mercedes"]
									force: 0.7
									objects: [
										{id: "moveto-uuid"},
										{beacon: "weaviate://localhost/other-moveto-uuid"},
									]
								}
							}
							) {
							beacon className
						}
			}`,
			expectedParamsToTraverser: traverser.ExploreParams{
				Limit: 17,
				NearText: &traverser.NearTextParams{
					Values: []string{"car", "best brand"},
					MoveTo: traverser.ExploreMove{
						Values: []string{"mercedes"},
						Force:  0.7,
						Objects: []traverser.ObjectMove{
							{ID: "moveto-uuid"},
							{Beacon: "weaviate://localhost/other-moveto-uuid"},
						},
					},
				},
			},
			resolverReturn: []search.Result{
				search.Result{
					Beacon:    "weaviate://localhost/some-uuid",
					ClassName: "bestClass",
				},
			},
			expectedResults: []result{{
				pathToField: []string{"Explore"},
				expectedValue: []interface{}{
					map[string]interface{}{
						"beacon":    "weaviate://localhost/some-uuid",
						"className": "bestClass",
					},
				},
			}},
		},

		testCase{
			name: "with moveTo and moveAwayFrom and objects set",
			query: `
			{
					Explore(
							limit: 17
							nearText: {
								concepts: ["car", "best brand"]
								moveTo: {
									concepts: ["mercedes"]
									force: 0.7
									objects: [
										{id: "moveto-uuid1"},
										{beacon: "weaviate://localhost/moveto-uuid2"},
									]
								}
								moveAwayFrom: {
									concepts: ["van"]
									force: 0.7
									objects: [
										{id: "moveAway-uuid1"},
										{beacon: "weaviate://localhost/moveAway-uuid2"},
										{id: "moveAway-uuid3"},
										{id: "moveAway-uuid4"},
									]
								}
							}
							) {
							beacon className
						}
			}`,
			expectedParamsToTraverser: traverser.ExploreParams{
				Limit: 17,
				NearText: &traverser.NearTextParams{
					Values: []string{"car", "best brand"},
					MoveTo: traverser.ExploreMove{
						Values: []string{"mercedes"},
						Force:  0.7,
						Objects: []traverser.ObjectMove{
							{ID: "moveto-uuid1"},
							{Beacon: "weaviate://localhost/moveto-uuid2"},
						},
					},
					MoveAwayFrom: traverser.ExploreMove{
						Values: []string{"van"},
						Force:  0.7,
						Objects: []traverser.ObjectMove{
							{ID: "moveAway-uuid1"},
							{Beacon: "weaviate://localhost/moveAway-uuid2"},
							{ID: "moveAway-uuid3"},
							{ID: "moveAway-uuid4"},
						},
					},
				},
			},
			resolverReturn: []search.Result{
				search.Result{
					Beacon:    "weaviate://localhost/some-uuid",
					ClassName: "bestClass",
				},
			},
			expectedResults: []result{{
				pathToField: []string{"Explore"},
				expectedValue: []interface{}{
					map[string]interface{}{
						"beacon":    "weaviate://localhost/some-uuid",
						"className": "bestClass",
					},
				},
			}},
		},
	}

	tests := testCases{
		testCase{
			name: "Resolve Explore with nearVector",
			query: `
			{
					Explore(nearVector: {vector: [0, 1, 0.8]}) {
							beacon className certainty
					}
			}`,
			expectedParamsToTraverser: traverser.ExploreParams{
				NearVector: &traverser.NearVectorParams{
					Vector: []float32{0, 1, 0.8},
				},
			},
			resolverReturn: []search.Result{
				search.Result{
					Beacon:    "weaviate://localhost/some-uuid",
					ClassName: "bestClass",
					Certainty: 0.7,
				},
			},
			expectedResults: []result{{
				pathToField: []string{"Explore"},
				expectedValue: []interface{}{
					map[string]interface{}{
						"beacon":    "weaviate://localhost/some-uuid",
						"className": "bestClass",
						"certainty": float32(0.7),
					},
				},
			}},
		},

		testCase{
			name: "with nearVector with optional limit",
			query: `
			{
					Explore(limit: 17, nearVector: {vector: [0, 1, 0.8]}) {
							beacon className certainty
					}
			}`,
			expectedParamsToTraverser: traverser.ExploreParams{
				NearVector: &traverser.NearVectorParams{
					Vector: []float32{0, 1, 0.8},
				},
				Limit: 17,
			},
			resolverReturn: []search.Result{
				search.Result{
					Beacon:    "weaviate://localhost/some-uuid",
					ClassName: "bestClass",
					Certainty: 0.7,
				},
			},
			expectedResults: []result{{
				pathToField: []string{"Explore"},
				expectedValue: []interface{}{
					map[string]interface{}{
						"beacon":    "weaviate://localhost/some-uuid",
						"className": "bestClass",
						"certainty": float32(0.7),
					},
				},
			}},
		},

		testCase{
			name: "Resolve Explore with nearObject and beacon set",
			query: `
			{
				Explore(
					nearObject: {
						beacon: "weaviate://localhost/27b5213d-e152-4fea-bd63-2063d529024d"
						certainty: 0.7
					}) {
						beacon className certainty
					}
			}`,
			expectedParamsToTraverser: traverser.ExploreParams{
				NearObject: &traverser.NearObjectParams{
					Beacon:    "weaviate://localhost/27b5213d-e152-4fea-bd63-2063d529024d",
					Certainty: 0.7,
				},
			},
			resolverReturn: []search.Result{
				{
					Beacon:    "weaviate://localhost/27b5213d-e152-4fea-bd63-2063d529024d",
					ClassName: "bestClass",
					Certainty: 0.7,
				},
			},
			expectedResults: []result{{
				pathToField: []string{"Explore"},
				expectedValue: []interface{}{
					map[string]interface{}{
						"beacon":    "weaviate://localhost/27b5213d-e152-4fea-bd63-2063d529024d",
						"className": "bestClass",
						"certainty": float32(0.7),
					},
				},
			}},
		},

		testCase{
			name: "Resolve Explore with nearObject and id set",
			query: `
			{
					Explore(
							limit: 17
							nearObject: {
								id: "27b5213d-e152-4fea-bd63-2063d529024d"
								certainty: 0.7
							}
							) {
							beacon className
						}
			}`,
			expectedParamsToTraverser: traverser.ExploreParams{
				Limit: 17,
				NearObject: &traverser.NearObjectParams{
					ID:        "27b5213d-e152-4fea-bd63-2063d529024d",
					Certainty: 0.7,
				},
			},
			resolverReturn: []search.Result{
				search.Result{
					Beacon:    "weaviate://localhost/some-uuid",
					ClassName: "bestClass",
				},
			},
			expectedResults: []result{{
				pathToField: []string{"Explore"},
				expectedValue: []interface{}{
					map[string]interface{}{
						"beacon":    "weaviate://localhost/some-uuid",
						"className": "bestClass",
					},
				},
			}},
		},
	}

	tests.AssertExtraction(t, newMockResolver())
	testsNearText.AssertExtraction(t, newMockResolver())
	tests.AssertExtraction(t, newMockResolverNoModules())
}

func Test_ExploreWithNoText2VecClasses(t *testing.T) {
	resolver := newMockResolverEmptySchema()
	query := `
	{
			Explore(
				nearText: {concepts: ["car", "best brand"], certainty: 0.6}, limit: 17 
				){
					beacon className
		}
	}`
	res := resolver.Resolve(query)
	require.Len(t, res.Errors, 1)
	assert.Contains(t, res.Errors[0].Message, "Unknown argument \"nearText\" on field \"Explore\"")
}

func Test_ExploreWithNoModules(t *testing.T) {
	resolver := newMockResolverNoModules()
	query := `
	{
			Explore(
				nearText: {concepts: ["car", "best brand"], certainty: 0.6}, limit: 17 
				){
					beacon className
		}
	}`
	res := resolver.Resolve(query)
	require.Len(t, res.Errors, 1)
	assert.Contains(t, res.Errors[0].Message, "Unknown argument \"nearText\" on field \"Explore\"")
}

func (tests testCases) AssertExtraction(t *testing.T, resolver *mockResolver) {
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			resolver.On("Explore", testCase.expectedParamsToTraverser).
				Return(testCase.resolverReturn, nil).Once()

			result := resolver.AssertResolve(t, testCase.query)

			for _, expectedResult := range testCase.expectedResults {
				value := result.Get(expectedResult.pathToField...).Result

				assert.Equal(t, expectedResult.expectedValue, value)
			}
		})
	}
}
