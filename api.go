package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vegadelalyra/go_jwt_authenticate/handler"
	"github.com/vegadelalyra/go_jwt_authenticate/middleware"
)

func main() {
	r := gin.Default()
	r.POST("/login", handler.LoginHandler)

	api := r.Group("/api")
	api.Use(middleware.ValidateToken())

	product := api.Group("/product")
	product.Use(middleware.Authorization([]int{1}))

	product.GET("/", handler.GetAll)
	product.POST("/", middleware.Authorization([]int{3}), handler.AddProduct)

	user := api.Group("/user")

	user.GET("/", func(c *gin.Context) {
		c.AbortWithStatusJSON(200, gin.H{
			"status": "ok",
		})
	})

	r.Run("localhost:8080")
}
