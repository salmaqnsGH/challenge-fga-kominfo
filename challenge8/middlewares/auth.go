package middlewares

import (
	database "latihan-jwt/database"
	"latihan-jwt/helpers"
	"latihan-jwt/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := helpers.VerifyToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   err.Error(),
			})
			return
		}

		c.Set("userData", claims)
		c.Next()
	}
}

func ProductAuthorizationPUT() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		product := models.Product{}

		if !isAdmin(c, userID) {
			productID, err := strconv.Atoi(c.Param("productID"))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "Unauthorized",
					"error":   "Invalid product ID data type",
				})
				return
			}

			err = db.Select("user_id").First(&product, uint(productID)).Error
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "Unauthorized",
					"error":   "Failed to find product",
				})
				return
			}

			if product.UserID != userID {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"message": "Forbidden",
					"error":   "You are not allowed to access this product",
				})
				return
			}
		}

		c.Next()
	}
}

func ProductAuthorizationDELETE() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		product := models.Product{}

		if !isAdmin(c, userID) {
			productID, err := strconv.Atoi(c.Param("productID"))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "Unauthorized",
					"error":   "Invalid product ID data type",
				})
				return
			}

			err = db.Select("user_id").First(&product, uint(productID)).Error
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "Unauthorized",
					"error":   "Failed to find product",
				})
				return
			}

			if product.UserID != userID {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"message": "Forbidden",
					"error":   "You are not allowed to access this product",
				})
				return
			}
		}

		c.Next()
	}
}

func ProductAuthorizationGET() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		product := models.Product{}

		if !isAdmin(c, userID) {
			productID, err := strconv.Atoi(c.Param("productID"))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "Unauthorized",
					"error":   "Invalid product ID data type",
				})
				return
			}

			err = db.Select("user_id").First(&product, uint(productID)).Error
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "Unauthorized",
					"error":   "Failed to find product",
				})
				return
			}

			if product.UserID != userID {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"message": "Forbidden",
					"error":   "You are not allowed to access this product",
				})
				return
			}
		}

		c.Next()
	}
}

func isAdmin(c *gin.Context, userID uint) bool {
	db := database.GetDB()
	user := models.User{}

	err := db.Select("role").First(&user, userID).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Unauthorized",
			"error":   "Failed to find product",
		})
		return false
	}

	if user.Role != "admin" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "Forbidden",
			"error":   "Access forbidden for this method",
		})
		return false
	}

	return true
}

func AuthorizeAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)

		userID := uint(userData["id"].(float64))
		user := models.User{}

		err := db.Select("role").First(&user, userID).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Unauthorized",
				"error":   "Failed to find product",
			})
			return
		}

		if user.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Forbidden",
				"error":   "Access forbidden for this method",
			})
			return
		}

		c.Next()
	}
}
