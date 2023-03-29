package controller

import (
	"belajar-gin/models"
	"belajar-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bookController struct {
	service service.Service
}

func NewBookController(service service.Service) *bookController {
	return &bookController{service}
}

func (h *bookController) CreateBook(ctx *gin.Context) {
	var input models.BookInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook := h.service.CreateBook(input)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, newBook)
}

// func (r *bookController) UpdateBook(ctx *gin.Context) {
// 	var bookInput models.Book
// 	bookID := ctx.Param("bookID")
// 	bookIDInt, err := strconv.Atoi(bookID)
// 	if err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 	}

// 	err = ctx.ShouldBindJSON(&bookInput)
// 	if err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 	}

// 	bookInput.ID = bookIDInt
// 	affectedRow, _ := r.repository.UpdateBook(bookInput)

// 	if affectedRow == 0 {
// 		ctx.JSON(http.StatusCreated, "ID not found")
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, "Updated")
// }

// func (r *bookController) GetBook(ctx *gin.Context) {
// 	bookID := ctx.Param("bookID")
// 	bookIDInt, err := strconv.Atoi(bookID)
// 	if err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 	}

// 	book, err := r.repository.GetBookByID(bookIDInt)
// 	if err != nil {
// 		ctx.JSON(http.StatusNotFound, "ID not found")
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, book)
// }

// func (r *bookController) GetAllBook(ctx *gin.Context) {
// 	books := r.repository.GetBooks()
// 	ctx.JSON(http.StatusOK, books)
// }

// func (r *bookController) DeleteBook(ctx *gin.Context) {
// 	bookID := ctx.Param("bookID")
// 	bookIDInt, err := strconv.Atoi(bookID)
// 	if err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 	}

// 	affectedRow := r.repository.DeleteBook(bookIDInt)

// 	if affectedRow == 0 {
// 		ctx.JSON(http.StatusCreated, "ID not found")
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, "Deleted")
// }
