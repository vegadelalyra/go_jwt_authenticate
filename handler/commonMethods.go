package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vegadelalyra/go_jwt_authenticate/models"
)

func ok(context *gin.Context, status int, message string, data interface{}) {
	context.AbortWithStatusJSON(http.StatusOK, models.Response{
		Data:    data,
		Status:  status,
		Message: message,
	})
}

func badRequest(context *gin.Context, status int, message string, errors []models.ErrorDetail) {
	context.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
		Error:   errors,
		Status:  status,
		Message: message,
	})
}
