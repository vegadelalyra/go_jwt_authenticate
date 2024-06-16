package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vegadelalyra/go_jwt_authenticate/models"
	"github.com/vegadelalyra/go_jwt_authenticate/token"
)

func ValidateToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.Request.Header.Get("apikey")
		referer := context.Request.Header.Get("Referer")
		valid, claims := token.VerifyToken(tokenString, referer)
		if !valid {
			ReturnUnauthorized(context)
		}
		if len(context.Keys) == 0 {
			context.Keys = make(map[string]interface{})
		}
		context.Keys["companyId"] = claims.CompanyID
		context.Keys["username"] = claims.Username
		context.Keys["roles"] = claims.Roles
	}
}

func ReturnUnauthorized(context *gin.Context) {
	context.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{
		Error: []models.ErrorDetail{
			{
				ErrorType:    models.ErrorTypeUnauthorized,
				ErrorMessage: "You are not authorized to access this path",
			},
		},
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized access",
	})
}

func Authorization(validRoles []int) gin.HandlerFunc {
	return func(context *gin.Context) {
		if len(context.Keys) == 0 {
			ReturnUnauthorized(context)
		}

		rolesVal := context.Keys["roles"]
		if rolesVal == nil {
			ReturnUnauthorized(context)
		}

		roles := rolesVal.([]int)
		validation := make(map[int]int)
		for _, val := range roles {
			validation[val] = 0
		}

		for _, val := range validRoles {
			if _, ok := validation[val]; !ok {
				ReturnUnauthorized(context)
			}
		}
	}
}
