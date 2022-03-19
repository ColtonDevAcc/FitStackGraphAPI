<<<<<<< HEAD
// Code generated by mockery v2.9.4. DO NOT EDIT.
=======
// Code generated by mockery v2.10.0. DO NOT EDIT.
>>>>>>> 05cda12d7a6807133a4cc9757955f753adf7ef09

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	graph "github.com/voodoostack/fitstackapi/graph"
)

// MutationResolver is an autogenerated mock type for the MutationResolver type
type MutationResolver struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, input
func (_m *MutationResolver) Login(ctx context.Context, input graph.LoginInput) (*graph.AuthResponse, error) {
	ret := _m.Called(ctx, input)

	var r0 *graph.AuthResponse
	if rf, ok := ret.Get(0).(func(context.Context, graph.LoginInput) *graph.AuthResponse); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*graph.AuthResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, graph.LoginInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, input
func (_m *MutationResolver) Register(ctx context.Context, input graph.RegisterInput) (*graph.AuthResponse, error) {
	ret := _m.Called(ctx, input)

	var r0 *graph.AuthResponse
	if rf, ok := ret.Get(0).(func(context.Context, graph.RegisterInput) *graph.AuthResponse); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*graph.AuthResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, graph.RegisterInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
