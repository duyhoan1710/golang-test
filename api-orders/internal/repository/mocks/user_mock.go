package repository_mock

import (
	"context"

	"api-orders/internal/model"

	mock "github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (or *UserRepositoryMock) Create(c context.Context, user *model.User) error {
	ret := or.Called(c, user)

	var r error

	if rf, ok := ret.Get(0).(func(context.Context, *model.User) error); ok {
		rf(c, user)
	} else {
		r = ret.Error(1)
	}

	return r
}

func (or *UserRepositoryMock) Find(c context.Context) ([]model.User, error) {
	ret := or.Called(c)

	var r1 []model.User

	if rf, ok := ret.Get(0).(func(context.Context) []model.User); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Get(0).([]model.User)
	}

	var r2 error

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r2 = rf(c)
	} else {
		r2 = ret.Error(1)
	}

	return r1, r2
}

func (or *UserRepositoryMock) FindByEmail(c context.Context, email string) (model.User, error) {
	ret := or.Called(c, email)

	var r1 model.User

	if rf, ok := ret.Get(0).(func(context.Context, string) model.User); ok {
		r1 = rf(c, email)
	} else {
		r1 = ret.Get(0).(model.User)
	}

	var r2 error

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r2 = rf(c, email)
	} else {
		r2 = ret.Error(1)
	}

	return r1, r2
}

func (or *UserRepositoryMock) FindById(c context.Context, id string) (model.User, error) {
	ret := or.Called(c, id)

	var r1 model.User

	if rf, ok := ret.Get(0).(func(context.Context, string) model.User); ok {
		r1 = rf(c, id)
	} else {
		r1 = ret.Get(0).(model.User)
	}

	var r2 error

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r2 = rf(c, id)
	} else {
		r2 = ret.Error(1)
	}

	return r1, r2
}
