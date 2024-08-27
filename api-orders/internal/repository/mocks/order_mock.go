package repository_mock

import (
	"context"

	"api-orders/internal/model"

	mock "github.com/stretchr/testify/mock"
)

type OrderRepositoryMock struct {
	mock.Mock
}

func (or *OrderRepositoryMock) Create(c context.Context, order *model.Order) error {
	ret := or.Called(c, order)

	var r error

	if rf, ok := ret.Get(0).(func(context.Context, *model.Order)); ok {
		rf(c, order)
	} else {
		r = ret.Error(0)
	}

	return r
}

func (or *OrderRepositoryMock) UpdateOne(c context.Context, orderId string, order *model.Order) error {
	ret := or.Called(c, order)

	var r error

	if rf, ok := ret.Get(0).(func(context.Context, string, *model.Order)); ok {
		rf(c, orderId, order)
	} else {
		r = ret.Error(0)
	}

	return r
}

func (or *OrderRepositoryMock) Find(c context.Context) ([]model.Order, error) {
	ret := or.Called(c)

	var r1 []model.Order

	if rf, ok := ret.Get(0).(func(context.Context) []model.Order); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Get(0).([]model.Order)
	}

	var r2 error

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r2 = rf(c)
	} else {
		r2 = ret.Error(1)
	}

	return r1, r2
}

func (or *OrderRepositoryMock) FindById(c context.Context, id string) (model.Order, error) {
	ret := or.Called(c, id)

	var r1 model.Order

	if rf, ok := ret.Get(0).(func(context.Context, string) model.Order); ok {
		r1 = rf(c, id)
	} else {
		r1 = ret.Get(0).(model.Order)
	}

	var r2 error

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r2 = rf(c, id)
	} else {
		r2 = ret.Error(1)
	}

	return r1, r2
}

func (or *OrderRepositoryMock) FindByUserId(c context.Context, userId string) ([]model.Order, error) {
	ret := or.Called(c, userId)

	var r1 []model.Order

	if rf, ok := ret.Get(0).(func(context.Context, string) []model.Order); ok {
		r1 = rf(c, userId)
	} else {
		r1 = ret.Get(0).([]model.Order)
	}

	var r2 error

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r2 = rf(c, userId)
	} else {
		r2 = ret.Error(1)
	}

	return r1, r2
}
