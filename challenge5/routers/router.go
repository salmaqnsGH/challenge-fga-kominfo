package router

import (
	controller "belajar-gin/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	r.POST("/books", controller.CreateBook)
	r.PUT("/books/:bookID", controller.UpdateBook)
	r.GET("/books/:bookID", controller.GetBook)
	r.DELETE("/books/:bookID", controller.DeleteBook)
	r.GET("/books", controller.GetAllBook)

	return r
}
