package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"api-orders/config"
	"api-orders/internal/exception"
	"api-orders/internal/model"
	"api-orders/internal/util"

	repositoryInterface "api-orders/internal/interface/repository"
	serviceInterface "api-orders/internal/interface/service"
)

type authService struct {
	UserRepository repositoryInterface.IUserRepository
	Env            *config.Env
}

func NewAuthService(userRepository repositoryInterface.IUserRepository, env *config.Env) serviceInterface.IAuthService {
	return &authService{
		UserRepository: userRepository,
		Env:            env,
	}
}

func (authService *authService) Login(c context.Context, email string, password string) (accessToken string, refreshToken string, customError exception.ICustomError) {
	user, err := authService.UserRepository.FindByEmail(c, email)
	if err != nil {
		return accessToken, refreshToken, exception.NewCustomError(exception.USER_NOT_FOUND)
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return accessToken, refreshToken, exception.NewCustomError(exception.PASSWORD_INCORRECT)
	}

	accessToken, err = util.CreateAccessToken(&user, authService.Env.AccessTokenSecret, authService.Env.AccessTokenExpiryHour)
	if err != nil {
		return accessToken, refreshToken, exception.NewCustomError(exception.INTERNAL_SERVER_ERROR, err.Error())
	}

	refreshToken, err = util.CreateRefreshToken(&user, authService.Env.RefreshTokenSecret, authService.Env.RefreshTokenExpiryHour)
	if err != nil {
		return accessToken, refreshToken, exception.NewCustomError(exception.INTERNAL_SERVER_ERROR, err.Error())
	}

	return accessToken, refreshToken, customError
}

func (authService *authService) Signup(c context.Context, name string, email string, password string) (accessToken string, refreshToken string, customError exception.ICustomError) {
	var err error = nil

	_, err = authService.UserRepository.FindByEmail(c, email)

	if err == nil {
		return accessToken, refreshToken, exception.NewCustomError(exception.USER_ALREADY_EXISTS)

	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return accessToken, refreshToken, exception.NewCustomError(exception.INTERNAL_SERVER_ERROR, err.Error())
	}

	user := model.User{
		ID:       primitive.NewObjectID(),
		Name:     name,
		Email:    email,
		Password: string(encryptedPassword),
	}

	err = authService.UserRepository.Create(c, &user)
	if err != nil {
		return accessToken, refreshToken, exception.NewCustomError(exception.INTERNAL_SERVER_ERROR, err.Error())
	}

	accessToken, err = util.CreateAccessToken(&user, authService.Env.AccessTokenSecret, authService.Env.AccessTokenExpiryHour)
	if err != nil {
		return accessToken, refreshToken, exception.NewCustomError(exception.INTERNAL_SERVER_ERROR, err.Error())
	}

	refreshToken, err = util.CreateRefreshToken(&user, authService.Env.RefreshTokenSecret, authService.Env.RefreshTokenExpiryHour)
	if err != nil {
		return accessToken, refreshToken, exception.NewCustomError(exception.INTERNAL_SERVER_ERROR, err.Error())
	}

	return accessToken, refreshToken, customError
}

func (authService *authService) RefreshToken(c context.Context, currentRefreshToken string) (accessToken string, refreshToken string, customError exception.ICustomError) {

	id, err := util.ExtractIDFromToken(currentRefreshToken, authService.Env.RefreshTokenSecret)
	if err != nil {
		return accessToken, refreshToken, exception.NewCustomError(exception.EXTRACT_TOKEN_ERROR)
	}

	user, err := authService.UserRepository.FindById(c, id)
	if err != nil {
		return accessToken, refreshToken, exception.NewCustomError(exception.USER_NOT_FOUND)
	}

	accessToken, err = util.CreateAccessToken(&user, authService.Env.AccessTokenSecret, authService.Env.AccessTokenExpiryHour)
	if err != nil {
		return accessToken, refreshToken, exception.NewCustomError(exception.INTERNAL_SERVER_ERROR, err.Error())
	}

	refreshToken, err = util.CreateRefreshToken(&user, authService.Env.RefreshTokenSecret, authService.Env.RefreshTokenExpiryHour)
	if err != nil {
		return accessToken, refreshToken, exception.NewCustomError(exception.INTERNAL_SERVER_ERROR, err.Error())
	}

	return accessToken, refreshToken, customError
}
