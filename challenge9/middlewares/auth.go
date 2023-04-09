package middlewares

import (
	database "latihan-jwt/database"
	"latihan-jwt/helpers"
	"latihan-jwt/models"
	"net/http"

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

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		if !isAdmin(c, userID) {
			if c.Request.Method == http.MethodDelete {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"message": "Forbidden",
					"error":   "Access forbidden for this method",
				})
				return
			} else if c.Request.Method == http.MethodPut {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"message": "Forbidden",
					"error":   "Access forbidden for this method",
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
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Not found",
			"error":   "Failed to find user",
		})
		return false
	}

	if user.Role != "admin" {
		return false
	}

	return true
}
