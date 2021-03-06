// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import context "context"

import mock "github.com/stretchr/testify/mock"
import runtime "k8s.io/apimachinery/pkg/runtime"
import types "k8s.io/apimachinery/pkg/types"

// RemoteEnvironmentManagerClient is an autogenerated mock type for the RemoteEnvironmentManagerClient type
type RemoteEnvironmentManagerClient struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, key, obj
func (_m *RemoteEnvironmentManagerClient) Get(ctx context.Context, key types.NamespacedName, obj runtime.Object) error {
	ret := _m.Called(ctx, key, obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.NamespacedName, runtime.Object) error); ok {
		r0 = rf(ctx, key, obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, obj
func (_m *RemoteEnvironmentManagerClient) Update(ctx context.Context, obj runtime.Object) error {
	ret := _m.Called(ctx, obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, runtime.Object) error); ok {
		r0 = rf(ctx, obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
