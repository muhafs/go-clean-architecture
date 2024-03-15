package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/muhafs/go-clean-architecture/domain/model/response"
	"github.com/muhafs/go-clean-architecture/helper/auth"
)

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		token := strings.Split(authHeader, " ")

		if len(token) != 2 {
			c.JSON(http.StatusUnauthorized, response.ErrorResponse{Success: false, Message: "Not authorized"})
			c.Abort()
		}

		tokenString := token[1]
		authorized, err := auth.Verify(tokenString, secret)
		if !authorized {
			c.JSON(http.StatusUnauthorized, response.ErrorResponse{Success: false, Message: err.Error()})
			c.Abort()
			return
		}

		userID, err := auth.ExtractID(tokenString, secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.ErrorResponse{Success: false, Message: err.Error()})
			c.Abort()
			return
		}

		c.Set("x-user-id", userID)
		c.Next()
	}
}
