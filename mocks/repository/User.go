// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"
	model "mock_gin/pkg/repository/model"

	mock "github.com/stretchr/testify/mock"
)

// User is an autogenerated mock type for the User type
type User struct {
	mock.Mock
}

// NewUser provides a mock function with given fields: ctx, name
func (_m *User) NewUser(ctx context.Context, name string) model.User {
	ret := _m.Called(ctx, name)

	var r0 model.User
	if rf, ok := ret.Get(0).(func(context.Context, string) model.User); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	return r0
}

// UpdateUser provides a mock function with given fields: ctx, id, name
func (_m *User) UpdateUser(ctx context.Context, id int64, name string) (model.User, error) {
	ret := _m.Called(ctx, id, name)

	var r0 model.User
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) model.User); ok {
		r0 = rf(ctx, id, name)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, string) error); ok {
		r1 = rf(ctx, id, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
