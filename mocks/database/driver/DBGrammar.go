// Code generated by mockery. DO NOT EDIT.

package driver

import (
	driver "github.com/goravel/framework/contracts/database/driver"
	mock "github.com/stretchr/testify/mock"

	squirrel "github.com/Masterminds/squirrel"
)

// DBGrammar is an autogenerated mock type for the DBGrammar type
type DBGrammar struct {
	mock.Mock
}

type DBGrammar_Expecter struct {
	mock *mock.Mock
}

func (_m *DBGrammar) EXPECT() *DBGrammar_Expecter {
	return &DBGrammar_Expecter{mock: &_m.Mock}
}

// CompileInRandomOrder provides a mock function with given fields: builder, conditions
func (_m *DBGrammar) CompileInRandomOrder(builder squirrel.SelectBuilder, conditions *driver.Conditions) squirrel.SelectBuilder {
	ret := _m.Called(builder, conditions)

	if len(ret) == 0 {
		panic("no return value specified for CompileInRandomOrder")
	}

	var r0 squirrel.SelectBuilder
	if rf, ok := ret.Get(0).(func(squirrel.SelectBuilder, *driver.Conditions) squirrel.SelectBuilder); ok {
		r0 = rf(builder, conditions)
	} else {
		r0 = ret.Get(0).(squirrel.SelectBuilder)
	}

	return r0
}

// DBGrammar_CompileInRandomOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CompileInRandomOrder'
type DBGrammar_CompileInRandomOrder_Call struct {
	*mock.Call
}

// CompileInRandomOrder is a helper method to define mock.On call
//   - builder squirrel.SelectBuilder
//   - conditions *driver.Conditions
func (_e *DBGrammar_Expecter) CompileInRandomOrder(builder interface{}, conditions interface{}) *DBGrammar_CompileInRandomOrder_Call {
	return &DBGrammar_CompileInRandomOrder_Call{Call: _e.mock.On("CompileInRandomOrder", builder, conditions)}
}

func (_c *DBGrammar_CompileInRandomOrder_Call) Run(run func(builder squirrel.SelectBuilder, conditions *driver.Conditions)) *DBGrammar_CompileInRandomOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(squirrel.SelectBuilder), args[1].(*driver.Conditions))
	})
	return _c
}

func (_c *DBGrammar_CompileInRandomOrder_Call) Return(_a0 squirrel.SelectBuilder) *DBGrammar_CompileInRandomOrder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DBGrammar_CompileInRandomOrder_Call) RunAndReturn(run func(squirrel.SelectBuilder, *driver.Conditions) squirrel.SelectBuilder) *DBGrammar_CompileInRandomOrder_Call {
	_c.Call.Return(run)
	return _c
}

// CompileLockForUpdate provides a mock function with given fields: builder, conditions
func (_m *DBGrammar) CompileLockForUpdate(builder squirrel.SelectBuilder, conditions *driver.Conditions) squirrel.SelectBuilder {
	ret := _m.Called(builder, conditions)

	if len(ret) == 0 {
		panic("no return value specified for CompileLockForUpdate")
	}

	var r0 squirrel.SelectBuilder
	if rf, ok := ret.Get(0).(func(squirrel.SelectBuilder, *driver.Conditions) squirrel.SelectBuilder); ok {
		r0 = rf(builder, conditions)
	} else {
		r0 = ret.Get(0).(squirrel.SelectBuilder)
	}

	return r0
}

// DBGrammar_CompileLockForUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CompileLockForUpdate'
type DBGrammar_CompileLockForUpdate_Call struct {
	*mock.Call
}

// CompileLockForUpdate is a helper method to define mock.On call
//   - builder squirrel.SelectBuilder
//   - conditions *driver.Conditions
func (_e *DBGrammar_Expecter) CompileLockForUpdate(builder interface{}, conditions interface{}) *DBGrammar_CompileLockForUpdate_Call {
	return &DBGrammar_CompileLockForUpdate_Call{Call: _e.mock.On("CompileLockForUpdate", builder, conditions)}
}

func (_c *DBGrammar_CompileLockForUpdate_Call) Run(run func(builder squirrel.SelectBuilder, conditions *driver.Conditions)) *DBGrammar_CompileLockForUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(squirrel.SelectBuilder), args[1].(*driver.Conditions))
	})
	return _c
}

func (_c *DBGrammar_CompileLockForUpdate_Call) Return(_a0 squirrel.SelectBuilder) *DBGrammar_CompileLockForUpdate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DBGrammar_CompileLockForUpdate_Call) RunAndReturn(run func(squirrel.SelectBuilder, *driver.Conditions) squirrel.SelectBuilder) *DBGrammar_CompileLockForUpdate_Call {
	_c.Call.Return(run)
	return _c
}

// CompileSharedLock provides a mock function with given fields: builder, conditions
func (_m *DBGrammar) CompileSharedLock(builder squirrel.SelectBuilder, conditions *driver.Conditions) squirrel.SelectBuilder {
	ret := _m.Called(builder, conditions)

	if len(ret) == 0 {
		panic("no return value specified for CompileSharedLock")
	}

	var r0 squirrel.SelectBuilder
	if rf, ok := ret.Get(0).(func(squirrel.SelectBuilder, *driver.Conditions) squirrel.SelectBuilder); ok {
		r0 = rf(builder, conditions)
	} else {
		r0 = ret.Get(0).(squirrel.SelectBuilder)
	}

	return r0
}

// DBGrammar_CompileSharedLock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CompileSharedLock'
type DBGrammar_CompileSharedLock_Call struct {
	*mock.Call
}

// CompileSharedLock is a helper method to define mock.On call
//   - builder squirrel.SelectBuilder
//   - conditions *driver.Conditions
func (_e *DBGrammar_Expecter) CompileSharedLock(builder interface{}, conditions interface{}) *DBGrammar_CompileSharedLock_Call {
	return &DBGrammar_CompileSharedLock_Call{Call: _e.mock.On("CompileSharedLock", builder, conditions)}
}

func (_c *DBGrammar_CompileSharedLock_Call) Run(run func(builder squirrel.SelectBuilder, conditions *driver.Conditions)) *DBGrammar_CompileSharedLock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(squirrel.SelectBuilder), args[1].(*driver.Conditions))
	})
	return _c
}

func (_c *DBGrammar_CompileSharedLock_Call) Return(_a0 squirrel.SelectBuilder) *DBGrammar_CompileSharedLock_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DBGrammar_CompileSharedLock_Call) RunAndReturn(run func(squirrel.SelectBuilder, *driver.Conditions) squirrel.SelectBuilder) *DBGrammar_CompileSharedLock_Call {
	_c.Call.Return(run)
	return _c
}

// NewDBGrammar creates a new instance of DBGrammar. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDBGrammar(t interface {
	mock.TestingT
	Cleanup(func())
}) *DBGrammar {
	mock := &DBGrammar{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
