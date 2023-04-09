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

	// productRouter := r.Group("products")
	// {
	// 	productRouter.Use(middlewares.Authentication())
	// 	productRouter.POST("/", controllers.CreateProduct)

	// 	productRouter.PUT("/:productID", middlewares.ProductAuthorizationPUT(), controllers.UpdateProduct)

	// 	productRouter.GET("/", controllers.GetProducts)
	// 	productRouter.GET("/:productID", middlewares.ProductAuthorizationGET(), controllers.GetProductByID)
	// 	productRouter.DELETE("/:productID", middlewares.ProductAuthorizationDELETE(), controllers.DeleteProductByID)
	// }

	productRepository := repositories.NeProductRepository(db)
	productService := services.NewProductService(productRepository)
	productController := controllers.NewProducrController(productService)
	// Define routes
	v1 := r.Group("/api/v1")
	{
		userRouter := r.Group("users")
		{
			userRouter.POST("/register", controllers.RegisterUser)
			userRouter.POST("/login", controllers.LoginUser)
		}
		productRouter := v1.Group("/products")
		{
			productRouter.Use(middlewares.Authentication())
			productRouter.POST("/", productController.CreateProduct)
			productRouter.GET("/", controllers.GetProducts)
		}
	}
	return r
}
