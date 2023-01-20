// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/megre/dto"

	mock "github.com/stretchr/testify/mock"
)

// CartServiceStore is an autogenerated mock type for the CartServiceStore type
type CartServiceStore struct {
	mock.Mock
}

// GetCartDetail provides a mock function with given fields: ctx, itemID, accountID
func (_m *CartServiceStore) GetCartDetail(ctx context.Context, itemID string, accountID string) (dto.Cart, error) {
	ret := _m.Called(ctx, itemID, accountID)

	var r0 dto.Cart
	if rf, ok := ret.Get(0).(func(context.Context, string, string) dto.Cart); ok {
		r0 = rf(ctx, itemID, accountID)
	} else {
		r0 = ret.Get(0).(dto.Cart)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, itemID, accountID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertCart provides a mock function with given fields: ctx, cartItem
func (_m *CartServiceStore) InsertCart(ctx context.Context, cartItem dto.Cart) (dto.Cart, error) {
	ret := _m.Called(ctx, cartItem)

	var r0 dto.Cart
	if rf, ok := ret.Get(0).(func(context.Context, dto.Cart) dto.Cart); ok {
		r0 = rf(ctx, cartItem)
	} else {
		r0 = ret.Get(0).(dto.Cart)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, dto.Cart) error); ok {
		r1 = rf(ctx, cartItem)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCartItems provides a mock function with given fields: ctx, accountID, pageNumber
func (_m *CartServiceStore) ListCartItems(ctx context.Context, accountID string, pageNumber int) ([]dto.CartItems, error) {
	ret := _m.Called(ctx, accountID, pageNumber)

	var r0 []dto.CartItems
	if rf, ok := ret.Get(0).(func(context.Context, string, int) []dto.CartItems); ok {
		r0 = rf(ctx, accountID, pageNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.CartItems)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int) error); ok {
		r1 = rf(ctx, accountID, pageNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateQuantity provides a mock function with given fields: ctx, cartID, quantity
func (_m *CartServiceStore) UpdateQuantity(ctx context.Context, cartID string, quantity int) (dto.Cart, error) {
	ret := _m.Called(ctx, cartID, quantity)

	var r0 dto.Cart
	if rf, ok := ret.Get(0).(func(context.Context, string, int) dto.Cart); ok {
		r0 = rf(ctx, cartID, quantity)
	} else {
		r0 = ret.Get(0).(dto.Cart)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int) error); ok {
		r1 = rf(ctx, cartID, quantity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCartServiceStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewCartServiceStore creates a new instance of CartServiceStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCartServiceStore(t mockConstructorTestingTNewCartServiceStore) *CartServiceStore {
	mock := &CartServiceStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}