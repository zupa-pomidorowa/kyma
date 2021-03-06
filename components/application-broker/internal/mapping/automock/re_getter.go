// Code generated by mockery v1.0.0
package automock

import internal "github.com/kyma-project/kyma/components/application-broker/internal"

import mock "github.com/stretchr/testify/mock"

// ReGetter is an autogenerated mock type for the ReGetter type
type ReGetter struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0
func (_m *ReGetter) Get(_a0 internal.RemoteEnvironmentName) (*internal.RemoteEnvironment, error) {
	ret := _m.Called(_a0)

	var r0 *internal.RemoteEnvironment
	if rf, ok := ret.Get(0).(func(internal.RemoteEnvironmentName) *internal.RemoteEnvironment); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internal.RemoteEnvironment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(internal.RemoteEnvironmentName) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
