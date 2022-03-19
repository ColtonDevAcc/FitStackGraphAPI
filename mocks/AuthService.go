<<<<<<< HEAD
// Code generated by mockery v2.9.4. DO NOT EDIT.
=======
// Code generated by mockery v2.10.0. DO NOT EDIT.
>>>>>>> 05cda12d7a6807133a4cc9757955f753adf7ef09

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	fitstackapi "github.com/voodoostack/fitstackapi"
)

// AuthService is an autogenerated mock type for the AuthService type
type AuthService struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, input
func (_m *AuthService) Login(ctx context.Context, input fitstackapi.LoginInput) (fitstackapi.AuthResponse, error) {
	ret := _m.Called(ctx, input)

	var r0 fitstackapi.AuthResponse
	if rf, ok := ret.Get(0).(func(context.Context, fitstackapi.LoginInput) fitstackapi.AuthResponse); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(fitstackapi.AuthResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, fitstackapi.LoginInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, input
func (_m *AuthService) Register(ctx context.Context, input fitstackapi.RegisterInput) (fitstackapi.AuthResponse, error) {
	ret := _m.Called(ctx, input)

	var r0 fitstackapi.AuthResponse
	if rf, ok := ret.Get(0).(func(context.Context, fitstackapi.RegisterInput) fitstackapi.AuthResponse); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(fitstackapi.AuthResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, fitstackapi.RegisterInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
