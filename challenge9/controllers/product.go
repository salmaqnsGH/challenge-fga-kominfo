package controllers

import (
	"errors"
	"latihan-jwt/models"
	"latihan-jwt/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type productController struct {
	productService services.ProductService
}

func NewProducrController(productService services.ProductService) *productController {
	return &productController{productService}
}

func (c *productController) CreateProduct(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	product := models.Product{}

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
			"error":    err.Error(),
		})
		return
	}

	product.UserID = uint(userData["id"].(float64))

	err = c.productService.CreateProduct(&product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func (c *productController) UpdateProduct(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	product := models.Product{}
	productID, _ := strconv.Atoi(ctx.Param("productID"))

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
			"error":    err.Error(),
		})
		return
	}

	product.UserID = uint(userData["id"].(float64))
	product.ID = uint(productID)

	updatedProduct, err := c.productService.UpdateProduct(&product)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"messsage": "Not found",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedProduct)
}

func (c *productController) GetProducts(ctx *gin.Context) {
	// userData := ctx.MustGet("userData").(jwt.MapClaims)
	// userID := uint(userData["id"].(float64))
	products := []models.Product{}

	products, err := c.productService.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (c *productController) GetProductByID(ctx *gin.Context) {
	productID, _ := strconv.Atoi(ctx.Param("productID"))

	product, err := c.productService.GetProductByID(uint(productID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"messsage": "Not found",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (c *productController) DeleteProductByID(ctx *gin.Context) {
	productID, err := strconv.Atoi(ctx.Param("productID"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Invalid product ID"))
		return
	}

	err = c.productService.DeleteProduct(uint(productID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"messsage": "Not found",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}
