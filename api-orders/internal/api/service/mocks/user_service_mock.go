package service_mock

import (
	"context"

	"github.com/stretchr/testify/mock"

	"api-orders/internal/exception"
	"api-orders/internal/model"
)

type UserServiceMock struct {
	mock.Mock
}

func (us *UserServiceMock) FindProfileById(c context.Context, userId string) (model.User, exception.ICustomError) {
	ret := us.Called(c, userId)

	var r1 model.User
	var r2 exception.ICustomError

	if rf, ok := ret.Get(0).(func(context.Context, string) model.User); ok {
		r1 = rf(c, userId)
	} else {
		r1 = ret.Get(0).(model.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) exception.ICustomError); ok {
		r2 = rf(c, userId)
	} else {
		r2 = ret.Get(1).(exception.ICustomError)
	}

	return r1, r2
}
