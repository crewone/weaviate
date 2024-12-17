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

// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	entitiesschema "github.com/weaviate/weaviate/entities/schema"

	models "github.com/weaviate/weaviate/entities/models"

	sharding "github.com/weaviate/weaviate/usecases/sharding"
)

// SchemaGetter is an autogenerated mock type for the SchemaGetter type
type SchemaGetter struct {
	mock.Mock
}

// ClusterHealthScore provides a mock function with given fields:
func (_m *SchemaGetter) ClusterHealthScore() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ClusterHealthScore")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// CopyShardingState provides a mock function with given fields: class
func (_m *SchemaGetter) CopyShardingState(class string) *sharding.State {
	ret := _m.Called(class)

	if len(ret) == 0 {
		panic("no return value specified for CopyShardingState")
	}

	var r0 *sharding.State
	if rf, ok := ret.Get(0).(func(string) *sharding.State); ok {
		r0 = rf(class)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sharding.State)
		}
	}

	return r0
}

// GetSchemaSkipAuth provides a mock function with given fields:
func (_m *SchemaGetter) GetSchemaSkipAuth() entitiesschema.Schema {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSchemaSkipAuth")
	}

	var r0 entitiesschema.Schema
	if rf, ok := ret.Get(0).(func() entitiesschema.Schema); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(entitiesschema.Schema)
	}

	return r0
}

// NodeName provides a mock function with given fields:
func (_m *SchemaGetter) NodeName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for NodeName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Nodes provides a mock function with given fields:
func (_m *SchemaGetter) Nodes() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Nodes")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// OptimisticTenantStatus provides a mock function with given fields: ctx, class, tenants
func (_m *SchemaGetter) OptimisticTenantStatus(ctx context.Context, class string, tenants string) (map[string]string, error) {
	ret := _m.Called(ctx, class, tenants)

	if len(ret) == 0 {
		panic("no return value specified for OptimisticTenantStatus")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (map[string]string, error)); ok {
		return rf(ctx, class, tenants)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) map[string]string); ok {
		r0 = rf(ctx, class, tenants)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, class, tenants)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadOnlyClass provides a mock function with given fields: _a0
func (_m *SchemaGetter) ReadOnlyClass(_a0 string) *models.Class {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for ReadOnlyClass")
	}

	var r0 *models.Class
	if rf, ok := ret.Get(0).(func(string) *models.Class); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Class)
		}
	}

	return r0
}

// ResolveParentNodes provides a mock function with given fields: _a0, _a1
func (_m *SchemaGetter) ResolveParentNodes(_a0 string, _a1 string) (map[string]string, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for ResolveParentNodes")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (map[string]string, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(string, string) map[string]string); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShardFromUUID provides a mock function with given fields: class, uuid
func (_m *SchemaGetter) ShardFromUUID(class string, uuid []byte) string {
	ret := _m.Called(class, uuid)

	if len(ret) == 0 {
		panic("no return value specified for ShardFromUUID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string, []byte) string); ok {
		r0 = rf(class, uuid)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ShardOwner provides a mock function with given fields: class, shard
func (_m *SchemaGetter) ShardOwner(class string, shard string) (string, error) {
	ret := _m.Called(class, shard)

	if len(ret) == 0 {
		panic("no return value specified for ShardOwner")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(class, shard)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(class, shard)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(class, shard)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShardReplicas provides a mock function with given fields: class, shard
func (_m *SchemaGetter) ShardReplicas(class string, shard string) ([]string, error) {
	ret := _m.Called(class, shard)

	if len(ret) == 0 {
		panic("no return value specified for ShardReplicas")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]string, error)); ok {
		return rf(class, shard)
	}
	if rf, ok := ret.Get(0).(func(string, string) []string); ok {
		r0 = rf(class, shard)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(class, shard)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Statistics provides a mock function with given fields:
func (_m *SchemaGetter) Statistics() map[string]interface{} {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Statistics")
	}

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func() map[string]interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// TenantsShards provides a mock function with given fields: ctx, class, tenants
func (_m *SchemaGetter) TenantsShards(ctx context.Context, class string, tenants ...string) (map[string]string, error) {
	_va := make([]interface{}, len(tenants))
	for _i := range tenants {
		_va[_i] = tenants[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, class)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for TenantsShards")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...string) (map[string]string, error)); ok {
		return rf(ctx, class, tenants...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...string) map[string]string); ok {
		r0 = rf(ctx, class, tenants...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...string) error); ok {
		r1 = rf(ctx, class, tenants...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSchemaGetter creates a new instance of SchemaGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSchemaGetter(t interface {
	mock.TestingT
	Cleanup(func())
},
) *SchemaGetter {
	mock := &SchemaGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
