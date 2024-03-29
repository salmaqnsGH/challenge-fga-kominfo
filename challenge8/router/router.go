package router

import (
	"latihan-jwt/controllers"
	"latihan-jwt/middlewares"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("users")
	{
		userRouter.POST("/register", controllers.RegisterUser)
		userRouter.POST("/login", controllers.LoginUser)
	}

	productRouter := r.Group("products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)

		productRouter.PUT("/:productID", middlewares.ProductAuthorizationPUT(), controllers.UpdateProduct)

		productRouter.GET("/", controllers.GetProducts)
		productRouter.GET("/:productID", middlewares.ProductAuthorizationGET(), controllers.GetProductByID)
		productRouter.DELETE("/:productID", middlewares.ProductAuthorizationDELETE(), controllers.DeleteProductByID)
	}

	return r
}
