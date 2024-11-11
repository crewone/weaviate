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

package authorization

// func Test_Authorizer(t *testing.T) {
// 	t.Run("when no authz is configured", func(t *testing.T) {
// 		cfg := config.Config{}

// 		logger, _ := test.NewNullLogger()
// 		authorizer := New(cfg, nil, logger)

// 		t.Run("it uses the dummy authorizer", func(t *testing.T) {
// 			_, ok := authorizer.(*DummyAuthorizer)
// 			assert.Equal(t, true, ok)
// 		})

// 		t.Run("any request is allowed", func(t *testing.T) {
// 			err := authorizer.Authorize(nil, DELETE, "the/world")
// 			assert.Nil(t, err)
// 		})
// 	})

// 	t.Run("when adminlist is configured", func(t *testing.T) {
// 		cfg := config.Config{
// 			Authorization: config.Authorization{
// 				AdminList: adminlist.Config{
// 					Enabled: true,
// 				},
// 			},
// 		}

// 		authorizer := adminlist.New(cfg.Authorization.AdminList)

// 		_, ok := authorizer.(*adminlist.Authorizer)
// 		assert.Equal(t, true, ok)
// 	})
// }
