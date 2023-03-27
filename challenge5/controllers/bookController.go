package controller

import (
	"belajar-gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateBook(ctx *gin.Context) {
	var newBook models.Book

	err := ctx.ShouldBindJSON(&newBook)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	newBook.ID = len(models.Books) + 1

	models.Books = append(models.Books, newBook)

	// ctx.JSON(http.StatusCreated, newBook)
	ctx.JSON(http.StatusCreated, "Created")
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

func GetAllBook(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, models.Books)
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	bookIndexToDelete := -1
	for i, book := range models.Books {
		if book.ID == bookIDInt {
			bookIndexToDelete = i
			break
		}
	}

	copy(models.Books[bookIndexToDelete:], models.Books[bookIndexToDelete+1:])
	models.Books[len(models.Books)-1] = models.Book{}
	models.Books = models.Books[:len(models.Books)-1]

	if bookIndexToDelete < 0 {
		ctx.JSON(http.StatusCreated, "ID not found")
		return
	}

	ctx.JSON(http.StatusCreated, "Deleted")
}
