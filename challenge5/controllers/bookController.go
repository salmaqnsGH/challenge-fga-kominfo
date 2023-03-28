package controller

import (
	"belajar-gin/models"
	"belajar-gin/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type bookController struct {
	repository repositories.Repository
}

func NewBookController(repository repositories.Repository) *bookController {
	return &bookController{repository}
}

func (r *bookController) CreateBook(ctx *gin.Context) {
	var newBook models.Book

	err := ctx.ShouldBindJSON(&newBook)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	r.repository.CreateBook(newBook)

	ctx.JSON(http.StatusCreated, "Created")
}

func (r *bookController) UpdateBook(ctx *gin.Context) {
	var bookInput models.Book
	bookID := ctx.Param("bookID")
	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	err = ctx.ShouldBindJSON(&bookInput)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	bookInput.ID = bookIDInt
	affectedRow, book := r.repository.UpdateBook(bookInput)

	if affectedRow == 0 {
		ctx.JSON(http.StatusCreated, "ID not found")
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func GetBook(ctx *gin.Context) {
	var b models.Book
	bookID := ctx.Param("bookID")
	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	isBookFound := false
	for _, book := range models.Books {
		if book.ID == bookIDInt {
			isBookFound = true
			b = book
			break
		}
	}

	if isBookFound {
		ctx.JSON(http.StatusCreated, b)
		return
	}

}

func (r *bookController) GetAllBook(ctx *gin.Context) {
	books := r.repository.GetBooks()
	ctx.JSON(http.StatusOK, books)
}

func (r *bookController) DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	affectedRow := r.repository.DeleteBook(bookIDInt)

	if affectedRow == 0 {
		ctx.JSON(http.StatusCreated, "ID not found")
		return
	}

	ctx.JSON(http.StatusOK, "Deleted")
}
