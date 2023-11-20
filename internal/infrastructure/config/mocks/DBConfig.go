// Code generated by mockery v2.36.0. DO NOT EDIT.

package config

import mock "github.com/stretchr/testify/mock"

// MockDBConfig is an autogenerated mock type for the DBConfig type
type MockDBConfig struct {
	mock.Mock
}

type MockDBConfig_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDBConfig) EXPECT() *MockDBConfig_Expecter {
	return &MockDBConfig_Expecter{mock: &_m.Mock}
}

// GetDatabaseName provides a mock function with given fields:
func (_m *MockDBConfig) GetDatabaseName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockDBConfig_GetDatabaseName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDatabaseName'
type MockDBConfig_GetDatabaseName_Call struct {
	*mock.Call
}

// GetDatabaseName is a helper method to define mock.On call
func (_e *MockDBConfig_Expecter) GetDatabaseName() *MockDBConfig_GetDatabaseName_Call {
	return &MockDBConfig_GetDatabaseName_Call{Call: _e.mock.On("GetDatabaseName")}
}

func (_c *MockDBConfig_GetDatabaseName_Call) Run(run func()) *MockDBConfig_GetDatabaseName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDBConfig_GetDatabaseName_Call) Return(_a0 string) *MockDBConfig_GetDatabaseName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDBConfig_GetDatabaseName_Call) RunAndReturn(run func() string) *MockDBConfig_GetDatabaseName_Call {
	_c.Call.Return(run)
	return _c
}

// GetHost provides a mock function with given fields:
func (_m *MockDBConfig) GetHost() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockDBConfig_GetHost_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHost'
type MockDBConfig_GetHost_Call struct {
	*mock.Call
}

// GetHost is a helper method to define mock.On call
func (_e *MockDBConfig_Expecter) GetHost() *MockDBConfig_GetHost_Call {
	return &MockDBConfig_GetHost_Call{Call: _e.mock.On("GetHost")}
}

func (_c *MockDBConfig_GetHost_Call) Run(run func()) *MockDBConfig_GetHost_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDBConfig_GetHost_Call) Return(_a0 string) *MockDBConfig_GetHost_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDBConfig_GetHost_Call) RunAndReturn(run func() string) *MockDBConfig_GetHost_Call {
	_c.Call.Return(run)
	return _c
}

// GetPassword provides a mock function with given fields:
func (_m *MockDBConfig) GetPassword() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockDBConfig_GetPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPassword'
type MockDBConfig_GetPassword_Call struct {
	*mock.Call
}

// GetPassword is a helper method to define mock.On call
func (_e *MockDBConfig_Expecter) GetPassword() *MockDBConfig_GetPassword_Call {
	return &MockDBConfig_GetPassword_Call{Call: _e.mock.On("GetPassword")}
}

func (_c *MockDBConfig_GetPassword_Call) Run(run func()) *MockDBConfig_GetPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDBConfig_GetPassword_Call) Return(_a0 string) *MockDBConfig_GetPassword_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDBConfig_GetPassword_Call) RunAndReturn(run func() string) *MockDBConfig_GetPassword_Call {
	_c.Call.Return(run)
	return _c
}

// GetPort provides a mock function with given fields:
func (_m *MockDBConfig) GetPort() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockDBConfig_GetPort_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPort'
type MockDBConfig_GetPort_Call struct {
	*mock.Call
}

// GetPort is a helper method to define mock.On call
func (_e *MockDBConfig_Expecter) GetPort() *MockDBConfig_GetPort_Call {
	return &MockDBConfig_GetPort_Call{Call: _e.mock.On("GetPort")}
}

func (_c *MockDBConfig_GetPort_Call) Run(run func()) *MockDBConfig_GetPort_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDBConfig_GetPort_Call) Return(_a0 int) *MockDBConfig_GetPort_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDBConfig_GetPort_Call) RunAndReturn(run func() int) *MockDBConfig_GetPort_Call {
	_c.Call.Return(run)
	return _c
}

// GetSSLMode provides a mock function with given fields:
func (_m *MockDBConfig) GetSSLMode() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockDBConfig_GetSSLMode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSSLMode'
type MockDBConfig_GetSSLMode_Call struct {
	*mock.Call
}

// GetSSLMode is a helper method to define mock.On call
func (_e *MockDBConfig_Expecter) GetSSLMode() *MockDBConfig_GetSSLMode_Call {
	return &MockDBConfig_GetSSLMode_Call{Call: _e.mock.On("GetSSLMode")}
}

func (_c *MockDBConfig_GetSSLMode_Call) Run(run func()) *MockDBConfig_GetSSLMode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDBConfig_GetSSLMode_Call) Return(_a0 bool) *MockDBConfig_GetSSLMode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDBConfig_GetSSLMode_Call) RunAndReturn(run func() bool) *MockDBConfig_GetSSLMode_Call {
	_c.Call.Return(run)
	return _c
}

// GetUser provides a mock function with given fields:
func (_m *MockDBConfig) GetUser() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockDBConfig_GetUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUser'
type MockDBConfig_GetUser_Call struct {
	*mock.Call
}

// GetUser is a helper method to define mock.On call
func (_e *MockDBConfig_Expecter) GetUser() *MockDBConfig_GetUser_Call {
	return &MockDBConfig_GetUser_Call{Call: _e.mock.On("GetUser")}
}

func (_c *MockDBConfig_GetUser_Call) Run(run func()) *MockDBConfig_GetUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDBConfig_GetUser_Call) Return(_a0 string) *MockDBConfig_GetUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDBConfig_GetUser_Call) RunAndReturn(run func() string) *MockDBConfig_GetUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDBConfig creates a new instance of MockDBConfig. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDBConfig(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDBConfig {
	mock := &MockDBConfig{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
