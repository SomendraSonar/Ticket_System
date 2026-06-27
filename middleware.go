package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(401, gin.H{
				"error": "Token required",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(
			authHeader,
			"Bearer ",
		)

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {
				return secretKey, nil
			},
		)

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		userID := int(claims["user_id"].(float64))

		c.Set("user_id", userID)

		c.Next()
	}
}
