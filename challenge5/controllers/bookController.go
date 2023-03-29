package controller

import (
	"belajar-gin/models"
	"belajar-gin/service"
	"net/http"
	"strconv"

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

func (h *bookController) GetBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book, err := h.service.GetBookyID(bookIDInt)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (h *bookController) UpdateBook(ctx *gin.Context) {
	var bookInput models.BookInput
	bookID := ctx.Param("bookID")
	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err = h.service.GetBookyID(bookIDInt)
	if err != nil {
		ctx.JSON(http.StatusOK, "ID not found")
		return
	}

	err = ctx.ShouldBindJSON(&bookInput)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updatedBook, err := h.service.UpdateBook(bookIDInt, bookInput)
	if err != nil {
		ctx.JSON(http.StatusOK, "ID not found")
		return
	}

	ctx.JSON(http.StatusOK, updatedBook)
}

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
