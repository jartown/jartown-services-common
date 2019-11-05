// Code generated by mockery v1.0.0. DO NOT EDIT.

package neomongo

import bson "go.mongodb.org/mongo-driver/bson"
import mock "github.com/stretchr/testify/mock"

// MockSingleResult is an autogenerated mock type for the SingleResult type
type MockSingleResult struct {
	mock.Mock
}

// Decode provides a mock function with given fields: v
func (_m *MockSingleResult) Decode(v interface{}) error {
	ret := _m.Called(v)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DecodeBytes provides a mock function with given fields:
func (_m *MockSingleResult) DecodeBytes() (bson.Raw, error) {
	ret := _m.Called()

	var r0 bson.Raw
	if rf, ok := ret.Get(0).(func() bson.Raw); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bson.Raw)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Err provides a mock function with given fields:
func (_m *MockSingleResult) Err() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}