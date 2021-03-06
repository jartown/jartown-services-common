// Code generated by mockery v1.0.0. DO NOT EDIT.

package neomongo

import context "context"
import mock "github.com/stretchr/testify/mock"

// MockCursor is an autogenerated mock type for the Cursor type
type MockCursor struct {
	mock.Mock
}

// All provides a mock function with given fields: ctx, results
func (_m *MockCursor) All(ctx context.Context, results interface{}) error {
	ret := _m.Called(ctx, results)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(ctx, results)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields: ctx
func (_m *MockCursor) Close(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Decode provides a mock function with given fields: val
func (_m *MockCursor) Decode(val interface{}) error {
	ret := _m.Called(val)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(val)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Err provides a mock function with given fields:
func (_m *MockCursor) Err() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *MockCursor) ID() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Next provides a mock function with given fields: ctx
func (_m *MockCursor) Next(ctx context.Context) bool {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
