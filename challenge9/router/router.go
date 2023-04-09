package router

import (
	"latihan-jwt/controllers"
	"latihan-jwt/middlewares"
	"latihan-jwt/repositories"
	"latihan-jwt/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	productRepository := repositories.NeProductRepository(db)
	productService := services.NewProductService(productRepository)
	productController := controllers.NewProducrController(productService)

	v1 := r.Group("/api/v1")
	{
		userRouter := v1.Group("users")
		{
			userRouter.POST("/register", controllers.RegisterUser)
			userRouter.POST("/login", controllers.LoginUser)
		}
		productRouter := v1.Group("/products")
		{
			productRouter.Use(middlewares.Authentication())

			productRouter.POST("/", productController.CreateProduct)
			productRouter.GET("/", productController.GetProducts)
			productRouter.DELETE("/:productID", middlewares.ProductAuthorization(), productController.DeleteProductByID)
			productRouter.PUT("/:productID", middlewares.ProductAuthorization(), productController.UpdateProduct)
			productRouter.GET("/:productID", productController.GetProductByID)
		}
	}
	return r
}
