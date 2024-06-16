package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vegadelalyra/go_jwt_authenticate/models"
)

type Product struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ProductRequest struct {
	Name string `json:"name"`
}

var db = []Product{
	{
		Id:   1,
		Name: "Product 1",
	},
	{
		Id:   2,
		Name: "Product 2",
	},
}

func GetAll(context *gin.Context) {
	context.AbortWithStatusJSON(http.StatusOK, db)
}

func AddProduct(context *gin.Context) {
	var productRequest ProductRequest
	if err := context.ShouldBindJSON(&productRequest); err != nil {
		var errors []models.ErrorDetail = make([]models.ErrorDetail, 0, 1)
		errors = append(errors, models.ErrorDetail{
			ErrorType:    models.ErrorTypeValidation,
			ErrorMessage: fmt.Sprintf("%v", err),
		})
		badRequest(context, http.StatusBadRequest, "invalid request", errors)
		return
	}

	product := Product{
		Id:   len(db) + 1,
		Name: productRequest.Name,
	}
	db = append(db, product)
	ok(context, http.StatusCreated, "Product Added", product)
}
