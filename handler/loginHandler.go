package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vegadelalyra/go_jwt_authenticate/models"
	"github.com/vegadelalyra/go_jwt_authenticate/token"
)

func LoginHandler(context *gin.Context) {
	var loginObj models.LoginRequest
	if err := context.ShouldBindJSON(&loginObj); err != nil {
		var errors []models.ErrorDetail = make([]models.ErrorDetail, 0, 1)
		errors = append(errors, models.ErrorDetail{
			ErrorType:    models.ErrorTypeValidation,
			ErrorMessage: fmt.Sprintf("%v", err),
		})
		badRequest(context, http.StatusBadRequest, "invalid request", errors)
	}

	var claims = &models.JwtClaims{}
	claims.CompanyID = "CompanyId"
	claims.Username = loginObj.UserName
	claims.Roles = []int{1, 2, 3}
	claims.Audience = context.Request.Header.Get("Referer")

	var tokenCreationTime = time.Now().UTC()
	var expirationTime = tokenCreationTime.Add(time.Duration(20) * time.Minute)
	tokenString, err := token.GenerateToken(claims, expirationTime)
	if err != nil {
		badRequest(context, http.StatusBadRequest, "error in generating token", []models.ErrorDetail{
			{
				ErrorType:    models.ErrorTypeError,
				ErrorMessage: err.Error(),
			},
		})
	}

	ok(context, http.StatusOK, "token created", tokenString)
}
