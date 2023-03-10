// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/megre/dto"

	mock "github.com/stretchr/testify/mock"
)

// AccountStore is an autogenerated mock type for the AccountStore type
type AccountStore struct {
	mock.Mock
}

// CreateAccount provides a mock function with given fields: ctx, request
func (_m *AccountStore) CreateAccount(ctx context.Context, request dto.CreateAccountRequest) (dto.Account, error) {
	ret := _m.Called(ctx, request)

	var r0 dto.Account
	if rf, ok := ret.Get(0).(func(context.Context, dto.CreateAccountRequest) dto.Account); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(dto.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, dto.CreateAccountRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountDetailsByEmail provides a mock function with given fields: ctx, email
func (_m *AccountStore) GetAccountDetailsByEmail(ctx context.Context, email string) (dto.Account, error) {
	ret := _m.Called(ctx, email)

	var r0 dto.Account
	if rf, ok := ret.Get(0).(func(context.Context, string) dto.Account); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(dto.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SuspendAccount provides a mock function with given fields: ctx, accountID
func (_m *AccountStore) SuspendAccount(ctx context.Context, accountID string) error {
	ret := _m.Called(ctx, accountID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, accountID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewAccountStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewAccountStore creates a new instance of AccountStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAccountStore(t mockConstructorTestingTNewAccountStore) *AccountStore {
	mock := &AccountStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
