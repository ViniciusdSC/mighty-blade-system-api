// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	context "context"

	clause "gorm.io/gorm/clause"

	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// GormDB is an autogenerated mock type for the GormDB type
type GormDB struct {
	mock.Mock
}

// AddError provides a mock function with given fields: err
func (_m *GormDB) AddError(err error) error {
	ret := _m.Called(err)

	var r0 error
	if rf, ok := ret.Get(0).(func(error) error); ok {
		r0 = rf(err)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Assign provides a mock function with given fields: attrs
func (_m *GormDB) Assign(attrs ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, attrs...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(...interface{}) *gorm.DB); ok {
		r0 = rf(attrs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Association provides a mock function with given fields: column
func (_m *GormDB) Association(column string) *gorm.Association {
	ret := _m.Called(column)

	var r0 *gorm.Association
	if rf, ok := ret.Get(0).(func(string) *gorm.Association); ok {
		r0 = rf(column)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.Association)
		}
	}

	return r0
}

// Attrs provides a mock function with given fields: attrs
func (_m *GormDB) Attrs(attrs ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, attrs...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(...interface{}) *gorm.DB); ok {
		r0 = rf(attrs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// AutoMigrate provides a mock function with given fields: dst
func (_m *GormDB) AutoMigrate(dst ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, dst...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(dst...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Begin provides a mock function with given fields: opts
func (_m *GormDB) Begin(opts ...*sql.TxOptions) *gorm.DB {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(...*sql.TxOptions) *gorm.DB); ok {
		r0 = rf(opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Clauses provides a mock function with given fields: conds
func (_m *GormDB) Clauses(conds ...clause.Expression) *gorm.DB {
	_va := make([]interface{}, len(conds))
	for _i := range conds {
		_va[_i] = conds[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(...clause.Expression) *gorm.DB); ok {
		r0 = rf(conds...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Commit provides a mock function with given fields:
func (_m *GormDB) Commit() *gorm.DB {
	ret := _m.Called()

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Connection provides a mock function with given fields: fc
func (_m *GormDB) Connection(fc func(*gorm.DB) error) error {
	ret := _m.Called(fc)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(*gorm.DB) error) error); ok {
		r0 = rf(fc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Count provides a mock function with given fields: count
func (_m *GormDB) Count(count *int64) *gorm.DB {
	ret := _m.Called(count)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(*int64) *gorm.DB); ok {
		r0 = rf(count)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Create provides a mock function with given fields: value
func (_m *GormDB) Create(value interface{}) *gorm.DB {
	ret := _m.Called(value)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}) *gorm.DB); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// CreateInBatches provides a mock function with given fields: value, batchSize
func (_m *GormDB) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	ret := _m.Called(value, batchSize)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, int) *gorm.DB); ok {
		r0 = rf(value, batchSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// DB provides a mock function with given fields:
func (_m *GormDB) DB() (*sql.DB, error) {
	ret := _m.Called()

	var r0 *sql.DB
	var r1 error
	if rf, ok := ret.Get(0).(func() (*sql.DB, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *sql.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.DB)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Debug provides a mock function with given fields:
func (_m *GormDB) Debug() *gorm.DB {
	ret := _m.Called()

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: value, conds
func (_m *GormDB) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, value)
	_ca = append(_ca, conds...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(value, conds...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Distinct provides a mock function with given fields: args
func (_m *GormDB) Distinct(args ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(...interface{}) *gorm.DB); ok {
		r0 = rf(args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Exec provides a mock function with given fields: _a0, values
func (_m *GormDB) Exec(_a0 string, values ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(string, ...interface{}) *gorm.DB); ok {
		r0 = rf(_a0, values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Find provides a mock function with given fields: dest, conds
func (_m *GormDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, dest)
	_ca = append(_ca, conds...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(dest, conds...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// FindInBatches provides a mock function with given fields: dest, batchSize, fc
func (_m *GormDB) FindInBatches(dest interface{}, batchSize int, fc func(*gorm.DB, int) error) *gorm.DB {
	ret := _m.Called(dest, batchSize, fc)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, int, func(*gorm.DB, int) error) *gorm.DB); ok {
		r0 = rf(dest, batchSize, fc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// First provides a mock function with given fields: dest, conds
func (_m *GormDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, dest)
	_ca = append(_ca, conds...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(dest, conds...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// FirstOrCreate provides a mock function with given fields: dest, conds
func (_m *GormDB) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, dest)
	_ca = append(_ca, conds...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(dest, conds...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// FirstOrInit provides a mock function with given fields: dest, conds
func (_m *GormDB) FirstOrInit(dest interface{}, conds ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, dest)
	_ca = append(_ca, conds...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(dest, conds...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Get provides a mock function with given fields: key
func (_m *GormDB) Get(key string) (interface{}, bool) {
	ret := _m.Called(key)

	var r0 interface{}
	var r1 bool
	if rf, ok := ret.Get(0).(func(string) (interface{}, bool)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Group provides a mock function with given fields: name
func (_m *GormDB) Group(name string) *gorm.DB {
	ret := _m.Called(name)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(string) *gorm.DB); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Having provides a mock function with given fields: query, args
func (_m *GormDB) Having(query interface{}, args ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// InnerJoins provides a mock function with given fields: query, args
func (_m *GormDB) InnerJoins(query string, args ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(string, ...interface{}) *gorm.DB); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// InstanceGet provides a mock function with given fields: key
func (_m *GormDB) InstanceGet(key string) (interface{}, bool) {
	ret := _m.Called(key)

	var r0 interface{}
	var r1 bool
	if rf, ok := ret.Get(0).(func(string) (interface{}, bool)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// InstanceSet provides a mock function with given fields: key, value
func (_m *GormDB) InstanceSet(key string, value interface{}) *gorm.DB {
	ret := _m.Called(key, value)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(string, interface{}) *gorm.DB); ok {
		r0 = rf(key, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Joins provides a mock function with given fields: query, args
func (_m *GormDB) Joins(query string, args ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(string, ...interface{}) *gorm.DB); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Last provides a mock function with given fields: dest, conds
func (_m *GormDB) Last(dest interface{}, conds ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, dest)
	_ca = append(_ca, conds...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(dest, conds...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Limit provides a mock function with given fields: limit
func (_m *GormDB) Limit(limit int) *gorm.DB {
	ret := _m.Called(limit)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(int) *gorm.DB); ok {
		r0 = rf(limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Migrator provides a mock function with given fields:
func (_m *GormDB) Migrator() gorm.Migrator {
	ret := _m.Called()

	var r0 gorm.Migrator
	if rf, ok := ret.Get(0).(func() gorm.Migrator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gorm.Migrator)
		}
	}

	return r0
}

// Model provides a mock function with given fields: value
func (_m *GormDB) Model(value interface{}) *gorm.DB {
	ret := _m.Called(value)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}) *gorm.DB); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Not provides a mock function with given fields: query, args
func (_m *GormDB) Not(query interface{}, args ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Offset provides a mock function with given fields: offset
func (_m *GormDB) Offset(offset int) *gorm.DB {
	ret := _m.Called(offset)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(int) *gorm.DB); ok {
		r0 = rf(offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Omit provides a mock function with given fields: columns
func (_m *GormDB) Omit(columns ...string) *gorm.DB {
	_va := make([]interface{}, len(columns))
	for _i := range columns {
		_va[_i] = columns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(...string) *gorm.DB); ok {
		r0 = rf(columns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Or provides a mock function with given fields: query, args
func (_m *GormDB) Or(query interface{}, args ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Order provides a mock function with given fields: value
func (_m *GormDB) Order(value interface{}) *gorm.DB {
	ret := _m.Called(value)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}) *gorm.DB); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Pluck provides a mock function with given fields: column, dest
func (_m *GormDB) Pluck(column string, dest interface{}) *gorm.DB {
	ret := _m.Called(column, dest)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(string, interface{}) *gorm.DB); ok {
		r0 = rf(column, dest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Preload provides a mock function with given fields: query, args
func (_m *GormDB) Preload(query string, args ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(string, ...interface{}) *gorm.DB); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Raw provides a mock function with given fields: _a0, values
func (_m *GormDB) Raw(_a0 string, values ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(string, ...interface{}) *gorm.DB); ok {
		r0 = rf(_a0, values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Rollback provides a mock function with given fields:
func (_m *GormDB) Rollback() *gorm.DB {
	ret := _m.Called()

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// RollbackTo provides a mock function with given fields: name
func (_m *GormDB) RollbackTo(name string) *gorm.DB {
	ret := _m.Called(name)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(string) *gorm.DB); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Row provides a mock function with given fields:
func (_m *GormDB) Row() *sql.Row {
	ret := _m.Called()

	var r0 *sql.Row
	if rf, ok := ret.Get(0).(func() *sql.Row); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Row)
		}
	}

	return r0
}

// Rows provides a mock function with given fields:
func (_m *GormDB) Rows() (*sql.Rows, error) {
	ret := _m.Called()

	var r0 *sql.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func() (*sql.Rows, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *sql.Rows); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: value
func (_m *GormDB) Save(value interface{}) *gorm.DB {
	ret := _m.Called(value)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}) *gorm.DB); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// SavePoint provides a mock function with given fields: name
func (_m *GormDB) SavePoint(name string) *gorm.DB {
	ret := _m.Called(name)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(string) *gorm.DB); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Scan provides a mock function with given fields: dest
func (_m *GormDB) Scan(dest interface{}) *gorm.DB {
	ret := _m.Called(dest)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}) *gorm.DB); ok {
		r0 = rf(dest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// ScanRows provides a mock function with given fields: rows, dest
func (_m *GormDB) ScanRows(rows *sql.Rows, dest interface{}) error {
	ret := _m.Called(rows, dest)

	var r0 error
	if rf, ok := ret.Get(0).(func(*sql.Rows, interface{}) error); ok {
		r0 = rf(rows, dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Scopes provides a mock function with given fields: funcs
func (_m *GormDB) Scopes(funcs ...func(*gorm.DB) *gorm.DB) *gorm.DB {
	_va := make([]interface{}, len(funcs))
	for _i := range funcs {
		_va[_i] = funcs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(...func(*gorm.DB) *gorm.DB) *gorm.DB); ok {
		r0 = rf(funcs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Select provides a mock function with given fields: query, args
func (_m *GormDB) Select(query interface{}, args ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Session provides a mock function with given fields: config
func (_m *GormDB) Session(config *gorm.Session) *gorm.DB {
	ret := _m.Called(config)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(*gorm.Session) *gorm.DB); ok {
		r0 = rf(config)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Set provides a mock function with given fields: key, value
func (_m *GormDB) Set(key string, value interface{}) *gorm.DB {
	ret := _m.Called(key, value)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(string, interface{}) *gorm.DB); ok {
		r0 = rf(key, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// SetupJoinTable provides a mock function with given fields: model, field, joinTable
func (_m *GormDB) SetupJoinTable(model interface{}, field string, joinTable interface{}) error {
	ret := _m.Called(model, field, joinTable)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, interface{}) error); ok {
		r0 = rf(model, field, joinTable)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Table provides a mock function with given fields: name, args
func (_m *GormDB) Table(name string, args ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, name)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(string, ...interface{}) *gorm.DB); ok {
		r0 = rf(name, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Take provides a mock function with given fields: dest, conds
func (_m *GormDB) Take(dest interface{}, conds ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, dest)
	_ca = append(_ca, conds...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(dest, conds...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// ToSQL provides a mock function with given fields: queryFn
func (_m *GormDB) ToSQL(queryFn func(*gorm.DB) *gorm.DB) string {
	ret := _m.Called(queryFn)

	var r0 string
	if rf, ok := ret.Get(0).(func(func(*gorm.DB) *gorm.DB) string); ok {
		r0 = rf(queryFn)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Transaction provides a mock function with given fields: fc, opts
func (_m *GormDB) Transaction(fc func(*gorm.DB) error, opts ...*sql.TxOptions) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, fc)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(*gorm.DB) error, ...*sql.TxOptions) error); ok {
		r0 = rf(fc, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unscoped provides a mock function with given fields:
func (_m *GormDB) Unscoped() *gorm.DB {
	ret := _m.Called()

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Update provides a mock function with given fields: column, value
func (_m *GormDB) Update(column string, value interface{}) *gorm.DB {
	ret := _m.Called(column, value)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(string, interface{}) *gorm.DB); ok {
		r0 = rf(column, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// UpdateColumn provides a mock function with given fields: column, value
func (_m *GormDB) UpdateColumn(column string, value interface{}) *gorm.DB {
	ret := _m.Called(column, value)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(string, interface{}) *gorm.DB); ok {
		r0 = rf(column, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// UpdateColumns provides a mock function with given fields: values
func (_m *GormDB) UpdateColumns(values interface{}) *gorm.DB {
	ret := _m.Called(values)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}) *gorm.DB); ok {
		r0 = rf(values)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Updates provides a mock function with given fields: values
func (_m *GormDB) Updates(values interface{}) *gorm.DB {
	ret := _m.Called(values)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}) *gorm.DB); ok {
		r0 = rf(values)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Use provides a mock function with given fields: plugin
func (_m *GormDB) Use(plugin gorm.Plugin) error {
	ret := _m.Called(plugin)

	var r0 error
	if rf, ok := ret.Get(0).(func(gorm.Plugin) error); ok {
		r0 = rf(plugin)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Where provides a mock function with given fields: query, args
func (_m *GormDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// WithContext provides a mock function with given fields: ctx
func (_m *GormDB) WithContext(ctx context.Context) *gorm.DB {
	ret := _m.Called(ctx)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(context.Context) *gorm.DB); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// NewGormDB creates a new instance of GormDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGormDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *GormDB {
	mock := &GormDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
