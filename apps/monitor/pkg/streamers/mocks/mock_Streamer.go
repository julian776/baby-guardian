// Code generated by mockery v2.43.2. DO NOT EDIT.

package streamers

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockStreamer is an autogenerated mock type for the Streamer type
type MockStreamer struct {
	mock.Mock
}

type MockStreamer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStreamer) EXPECT() *MockStreamer_Expecter {
	return &MockStreamer_Expecter{mock: &_m.Mock}
}

// GetRecords provides a mock function with given fields: ctx
func (_m *MockStreamer) GetRecords(ctx context.Context) ([][]byte, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetRecords")
	}

	var r0 [][]byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([][]byte, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) [][]byte); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStreamer_GetRecords_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRecords'
type MockStreamer_GetRecords_Call struct {
	*mock.Call
}

// GetRecords is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockStreamer_Expecter) GetRecords(ctx interface{}) *MockStreamer_GetRecords_Call {
	return &MockStreamer_GetRecords_Call{Call: _e.mock.On("GetRecords", ctx)}
}

func (_c *MockStreamer_GetRecords_Call) Run(run func(ctx context.Context)) *MockStreamer_GetRecords_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockStreamer_GetRecords_Call) Return(_a0 [][]byte, _a1 error) *MockStreamer_GetRecords_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStreamer_GetRecords_Call) RunAndReturn(run func(context.Context) ([][]byte, error)) *MockStreamer_GetRecords_Call {
	_c.Call.Return(run)
	return _c
}

// PutRecord provides a mock function with given fields: ctx, data
func (_m *MockStreamer) PutRecord(ctx context.Context, data []byte) error {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for PutRecord")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStreamer_PutRecord_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutRecord'
type MockStreamer_PutRecord_Call struct {
	*mock.Call
}

// PutRecord is a helper method to define mock.On call
//   - ctx context.Context
//   - data []byte
func (_e *MockStreamer_Expecter) PutRecord(ctx interface{}, data interface{}) *MockStreamer_PutRecord_Call {
	return &MockStreamer_PutRecord_Call{Call: _e.mock.On("PutRecord", ctx, data)}
}

func (_c *MockStreamer_PutRecord_Call) Run(run func(ctx context.Context, data []byte)) *MockStreamer_PutRecord_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]byte))
	})
	return _c
}

func (_c *MockStreamer_PutRecord_Call) Return(_a0 error) *MockStreamer_PutRecord_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStreamer_PutRecord_Call) RunAndReturn(run func(context.Context, []byte) error) *MockStreamer_PutRecord_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStreamer creates a new instance of MockStreamer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStreamer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStreamer {
	mock := &MockStreamer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
