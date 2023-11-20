// Code generated by mockery v2.36.0. DO NOT EDIT.

package services

import (
	models "github.com/ViniciusdSC/mighty-blade-system-api/internal/api/models"
	services "github.com/ViniciusdSC/mighty-blade-system-api/internal/api/services"
	mock "github.com/stretchr/testify/mock"
)

// MockItemService is an autogenerated mock type for the ItemService type
type MockItemService struct {
	mock.Mock
}

type MockItemService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockItemService) EXPECT() *MockItemService_Expecter {
	return &MockItemService_Expecter{mock: &_m.Mock}
}

// Count provides a mock function with given fields: filter
func (_m *MockItemService) Count(filter services.ItemFilter) (*int64, error) {
	ret := _m.Called(filter)

	var r0 *int64
	var r1 error
	if rf, ok := ret.Get(0).(func(services.ItemFilter) (*int64, error)); ok {
		return rf(filter)
	}
	if rf, ok := ret.Get(0).(func(services.ItemFilter) *int64); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int64)
		}
	}

	if rf, ok := ret.Get(1).(func(services.ItemFilter) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockItemService_Count_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Count'
type MockItemService_Count_Call struct {
	*mock.Call
}

// Count is a helper method to define mock.On call
//   - filter services.ItemFilter
func (_e *MockItemService_Expecter) Count(filter interface{}) *MockItemService_Count_Call {
	return &MockItemService_Count_Call{Call: _e.mock.On("Count", filter)}
}

func (_c *MockItemService_Count_Call) Run(run func(filter services.ItemFilter)) *MockItemService_Count_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(services.ItemFilter))
	})
	return _c
}

func (_c *MockItemService_Count_Call) Return(_a0 *int64, _a1 error) *MockItemService_Count_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockItemService_Count_Call) RunAndReturn(run func(services.ItemFilter) (*int64, error)) *MockItemService_Count_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: model
func (_m *MockItemService) Create(model *models.ItemModel) error {
	ret := _m.Called(model)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.ItemModel) error); ok {
		r0 = rf(model)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockItemService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockItemService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - model *models.ItemModel
func (_e *MockItemService_Expecter) Create(model interface{}) *MockItemService_Create_Call {
	return &MockItemService_Create_Call{Call: _e.mock.On("Create", model)}
}

func (_c *MockItemService_Create_Call) Run(run func(model *models.ItemModel)) *MockItemService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.ItemModel))
	})
	return _c
}

func (_c *MockItemService_Create_Call) Return(_a0 error) *MockItemService_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockItemService_Create_Call) RunAndReturn(run func(*models.ItemModel) error) *MockItemService_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: model
func (_m *MockItemService) Delete(model *models.ItemModel) error {
	ret := _m.Called(model)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.ItemModel) error); ok {
		r0 = rf(model)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockItemService_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockItemService_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - model *models.ItemModel
func (_e *MockItemService_Expecter) Delete(model interface{}) *MockItemService_Delete_Call {
	return &MockItemService_Delete_Call{Call: _e.mock.On("Delete", model)}
}

func (_c *MockItemService_Delete_Call) Run(run func(model *models.ItemModel)) *MockItemService_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.ItemModel))
	})
	return _c
}

func (_c *MockItemService_Delete_Call) Return(_a0 error) *MockItemService_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockItemService_Delete_Call) RunAndReturn(run func(*models.ItemModel) error) *MockItemService_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Find provides a mock function with given fields: filter
func (_m *MockItemService) Find(filter services.ItemFilter) ([]models.ItemModel, error) {
	ret := _m.Called(filter)

	var r0 []models.ItemModel
	var r1 error
	if rf, ok := ret.Get(0).(func(services.ItemFilter) ([]models.ItemModel, error)); ok {
		return rf(filter)
	}
	if rf, ok := ret.Get(0).(func(services.ItemFilter) []models.ItemModel); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ItemModel)
		}
	}

	if rf, ok := ret.Get(1).(func(services.ItemFilter) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockItemService_Find_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Find'
type MockItemService_Find_Call struct {
	*mock.Call
}

// Find is a helper method to define mock.On call
//   - filter services.ItemFilter
func (_e *MockItemService_Expecter) Find(filter interface{}) *MockItemService_Find_Call {
	return &MockItemService_Find_Call{Call: _e.mock.On("Find", filter)}
}

func (_c *MockItemService_Find_Call) Run(run func(filter services.ItemFilter)) *MockItemService_Find_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(services.ItemFilter))
	})
	return _c
}

func (_c *MockItemService_Find_Call) Return(_a0 []models.ItemModel, _a1 error) *MockItemService_Find_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockItemService_Find_Call) RunAndReturn(run func(services.ItemFilter) ([]models.ItemModel, error)) *MockItemService_Find_Call {
	_c.Call.Return(run)
	return _c
}

// FindOne provides a mock function with given fields: id
func (_m *MockItemService) FindOne(id string) (*models.ItemModel, error) {
	ret := _m.Called(id)

	var r0 *models.ItemModel
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*models.ItemModel, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *models.ItemModel); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ItemModel)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockItemService_FindOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindOne'
type MockItemService_FindOne_Call struct {
	*mock.Call
}

// FindOne is a helper method to define mock.On call
//   - id string
func (_e *MockItemService_Expecter) FindOne(id interface{}) *MockItemService_FindOne_Call {
	return &MockItemService_FindOne_Call{Call: _e.mock.On("FindOne", id)}
}

func (_c *MockItemService_FindOne_Call) Run(run func(id string)) *MockItemService_FindOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockItemService_FindOne_Call) Return(_a0 *models.ItemModel, _a1 error) *MockItemService_FindOne_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockItemService_FindOne_Call) RunAndReturn(run func(string) (*models.ItemModel, error)) *MockItemService_FindOne_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: model, values
func (_m *MockItemService) Update(model *models.ItemModel, values *models.ItemModel) error {
	ret := _m.Called(model, values)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.ItemModel, *models.ItemModel) error); ok {
		r0 = rf(model, values)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockItemService_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockItemService_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - model *models.ItemModel
//   - values *models.ItemModel
func (_e *MockItemService_Expecter) Update(model interface{}, values interface{}) *MockItemService_Update_Call {
	return &MockItemService_Update_Call{Call: _e.mock.On("Update", model, values)}
}

func (_c *MockItemService_Update_Call) Run(run func(model *models.ItemModel, values *models.ItemModel)) *MockItemService_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.ItemModel), args[1].(*models.ItemModel))
	})
	return _c
}

func (_c *MockItemService_Update_Call) Return(_a0 error) *MockItemService_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockItemService_Update_Call) RunAndReturn(run func(*models.ItemModel, *models.ItemModel) error) *MockItemService_Update_Call {
	_c.Call.Return(run)
	return _c
}

// Validate provides a mock function with given fields: model, ignoreRules
func (_m *MockItemService) Validate(model *models.ItemModel, ignoreRules ...string) error {
	_va := make([]interface{}, len(ignoreRules))
	for _i := range ignoreRules {
		_va[_i] = ignoreRules[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, model)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.ItemModel, ...string) error); ok {
		r0 = rf(model, ignoreRules...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockItemService_Validate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Validate'
type MockItemService_Validate_Call struct {
	*mock.Call
}

// Validate is a helper method to define mock.On call
//   - model *models.ItemModel
//   - ignoreRules ...string
func (_e *MockItemService_Expecter) Validate(model interface{}, ignoreRules ...interface{}) *MockItemService_Validate_Call {
	return &MockItemService_Validate_Call{Call: _e.mock.On("Validate",
		append([]interface{}{model}, ignoreRules...)...)}
}

func (_c *MockItemService_Validate_Call) Run(run func(model *models.ItemModel, ignoreRules ...string)) *MockItemService_Validate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(*models.ItemModel), variadicArgs...)
	})
	return _c
}

func (_c *MockItemService_Validate_Call) Return(_a0 error) *MockItemService_Validate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockItemService_Validate_Call) RunAndReturn(run func(*models.ItemModel, ...string) error) *MockItemService_Validate_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockItemService creates a new instance of MockItemService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockItemService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockItemService {
	mock := &MockItemService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
