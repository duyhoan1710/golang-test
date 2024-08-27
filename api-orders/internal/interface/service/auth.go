package service

import (
	"api-orders/internal/exception"
	"context"
)

type IAuthService interface {
	Login(c context.Context, email string, password string) (accessToken string, refreshToken string, customError exception.ICustomError)
	Signup(c context.Context, name string, email string, password string) (accessToken string, refreshToken string, customError exception.ICustomError)
	RefreshToken(c context.Context, currentRefreshToken string) (accessToken string, refreshToken string, customError exception.ICustomError)
}
