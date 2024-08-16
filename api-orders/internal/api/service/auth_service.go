package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"api-orders/config"
	"api-orders/internal/dto"
	model "api-orders/internal/model"
	"api-orders/internal/repository"
	"api-orders/internal/util"
)

type AuthService struct {
	UserRepository *repository.UserRepository
	Env            *config.Env
}

func (authService *AuthService) Login(c *gin.Context, email string, password string) (accessToken string, refreshToken string) {
	user, err := authService.UserRepository.FindByEmail(c, email)

	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Message: "User not found with the given email"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	accessToken, err = util.CreateAccessToken(&user, authService.Env.AccessTokenSecret, authService.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err = util.CreateRefreshToken(&user, authService.Env.RefreshTokenSecret, authService.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	return accessToken, refreshToken
}

func (authService *AuthService) Signup(c *gin.Context, name string, email string, password string) (accessToken string, refreshToken string) {
	var err error = nil

	_, err = authService.UserRepository.FindByEmail(c, email)

	if err == nil {
		c.JSON(http.StatusConflict, dto.ErrorResponse{Message: "User already exists with the given email"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	user := model.User{
		ID:       primitive.NewObjectID(),
		Name:     name,
		Email:    email,
		Password: string(encryptedPassword),
	}

	err = authService.UserRepository.Create(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, err = util.CreateAccessToken(&user, authService.Env.AccessTokenSecret, authService.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err = util.CreateRefreshToken(&user, authService.Env.RefreshTokenSecret, authService.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	return accessToken, refreshToken
}

func (authService *AuthService) RefreshToken(c *gin.Context, currentRefreshToken string) (accessToken string, refreshToken string) {

	id, err := util.ExtractIDFromToken(currentRefreshToken, authService.Env.RefreshTokenSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "User not found"})
		return
	}

	user, err := authService.UserRepository.FindById(c, id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "User not found"})
		return
	}

	accessToken, err = util.CreateAccessToken(&user, authService.Env.AccessTokenSecret, authService.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err = util.CreateRefreshToken(&user, authService.Env.RefreshTokenSecret, authService.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	return accessToken, refreshToken
}
