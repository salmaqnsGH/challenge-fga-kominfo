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

func (c *bookController) CreateBook(ctx *gin.Context) {
	var input models.BookInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook := c.service.CreateBook(input)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, newBook)
}

func (c *bookController) GetBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book, err := c.service.GetBookyID(bookIDInt)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "ID not found"})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (c *bookController) UpdateBook(ctx *gin.Context) {
	var bookInput models.BookInput
	bookID := ctx.Param("bookID")
	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err = c.service.GetBookyID(bookIDInt)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "ID not found"})
		return
	}

	err = ctx.ShouldBindJSON(&bookInput)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updatedBook, err := c.service.UpdateBook(bookIDInt, bookInput)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "ID not found"})
		return
	}

	ctx.JSON(http.StatusOK, updatedBook)
}

func (c *bookController) GetAllBook(ctx *gin.Context) {
	books, err := c.service.GetBooks()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (c *bookController) DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	err = c.service.DeleteBook(bookIDInt)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "ID not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Book deleted succesfully"})
}
