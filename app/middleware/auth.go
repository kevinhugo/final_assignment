package middleware

import (
	"net/http"

	"assignment/app/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := utils.TokenValid(c)
		if userId == 0 {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Set("UserID", userId)
		c.Next()
	}
}
