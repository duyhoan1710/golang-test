package grpc_mock

import "github.com/stretchr/testify/mock"

type PaymentGRPCClientMock struct {
	mock.Mock
}

func (s *PaymentGRPCClientMock) ProcessPayment(orderId string, userId string) bool {
	ret := s.Called(orderId, userId)

	var r bool

	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r = rf(orderId, userId)
	} else {
		r = ret.Get(0).(bool)
	}

	return r
}
