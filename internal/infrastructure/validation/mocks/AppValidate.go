// Code generated by mockery v2.36.0. DO NOT EDIT.

package validation

import mock "github.com/stretchr/testify/mock"

// MockAppValidate is an autogenerated mock type for the AppValidate type
type MockAppValidate struct {
	mock.Mock
}

type MockAppValidate_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAppValidate) EXPECT() *MockAppValidate_Expecter {
	return &MockAppValidate_Expecter{mock: &_m.Mock}
}

// StructExceptRules provides a mock function with given fields: s, rules
func (_m *MockAppValidate) StructExceptRules(s interface{}, rules ...string) error {
	_va := make([]interface{}, len(rules))
	for _i := range rules {
		_va[_i] = rules[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, s)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, ...string) error); ok {
		r0 = rf(s, rules...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAppValidate_StructExceptRules_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StructExceptRules'
type MockAppValidate_StructExceptRules_Call struct {
	*mock.Call
}

// StructExceptRules is a helper method to define mock.On call
//   - s interface{}
//   - rules ...string
func (_e *MockAppValidate_Expecter) StructExceptRules(s interface{}, rules ...interface{}) *MockAppValidate_StructExceptRules_Call {
	return &MockAppValidate_StructExceptRules_Call{Call: _e.mock.On("StructExceptRules",
		append([]interface{}{s}, rules...)...)}
}

func (_c *MockAppValidate_StructExceptRules_Call) Run(run func(s interface{}, rules ...string)) *MockAppValidate_StructExceptRules_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(interface{}), variadicArgs...)
	})
	return _c
}

func (_c *MockAppValidate_StructExceptRules_Call) Return(_a0 error) *MockAppValidate_StructExceptRules_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAppValidate_StructExceptRules_Call) RunAndReturn(run func(interface{}, ...string) error) *MockAppValidate_StructExceptRules_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAppValidate creates a new instance of MockAppValidate. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAppValidate(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAppValidate {
	mock := &MockAppValidate{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
