package middleware

import (
	"net/http"
	"strings"

	"api-payments/internal/dto"
	"api-payments/internal/util"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := util.IsAuthorized(authToken, secret)
			if authorized {
				userId, err := util.ExtractIDFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: err.Error()})
					c.Abort()
					return
				}
				c.Set("x-user-id", userId)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "Not authorized"})
		c.Abort()
	}
}
