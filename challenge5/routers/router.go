package router

import (
	controller "belajar-gin/controllers"
	"belajar-gin/repositories"
	"belajar-gin/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	bookRepository := repositories.NewRepository(db)
	bookService := service.NewService(bookRepository)
	bookController := controller.NewBookController(bookService)

	router.POST("/books", bookController.CreateBook)
	// router.PUT("/books/:bookID", bookController.UpdateBook)
	router.GET("/books/:bookID", bookController.GetBook)
	// router.DELETE("/books/:bookID", bookController.DeleteBook)
	// router.GET("/books", bookController.GetAllBook)

	return router
}
