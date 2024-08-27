package middleware

import (
	"api-orders/internal/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InternalServerErrorHandler(c *gin.Context) {
	c.Next()

	c.JSON(http.StatusInternalServerError, exception.NewCustomError(exception.INTERNAL_SERVER_ERROR, c.Errors.Errors()...))
}
