package router

import (
	controller "belajar-gin/controllers"
	"belajar-gin/repositories"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func StartServer(db *sql.DB) *gin.Engine {
	router := gin.Default()

	activityRepository := repositories.NewRepository(db)
	activityController := controller.NewBookController(activityRepository)

	router.POST("/books", activityController.CreateBook)
	router.PUT("/books/:bookID", controller.UpdateBook)
	router.GET("/books/:bookID", controller.GetBook)
	router.DELETE("/books/:bookID", controller.DeleteBook)
	router.GET("/books", controller.GetAllBook)

	return router
}
