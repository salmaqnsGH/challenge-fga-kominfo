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

	// ctx.JSON(http.StatusOK, book)
	ctx.JSON(http.StatusOK, "Created")
}

func UpdateBook(ctx *gin.Context) {
	var updatedBook models.Book
	bookID := ctx.Param("bookID")
	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	err = ctx.ShouldBindJSON(&updatedBook)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	isBookFound := false
	for i, book := range models.Books {
		if book.ID == bookIDInt {
			isBookFound = true
			models.Books[i] = updatedBook
			models.Books[i].ID = bookIDInt
			break
		}
	}

	if isBookFound {
		ctx.JSON(http.StatusCreated, "Updated")
		return
	}

	ctx.JSON(http.StatusCreated, "ID is not valid")
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
