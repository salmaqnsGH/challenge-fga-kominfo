package router

import (
	controller "belajar-gin/controllers"
	"belajar-gin/repositories"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func StartServer(db *sql.DB) *gin.Engine {
	router := gin.Default()

	bookRepository := repositories.NewRepository(db)
	bookController := controller.NewBookController(bookRepository)

	router.POST("/books", bookController.CreateBook)
	router.PUT("/books/:bookID", controller.UpdateBook)
	router.GET("/books/:bookID", controller.GetBook)
	router.DELETE("/books/:bookID", controller.DeleteBook)
	router.GET("/books", bookController.GetAllBook)

	return router
}
