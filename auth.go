package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("mysecretkey")

func Register(c *gin.Context) {
	var user User

	// Parse request body
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid data",
		})
		return
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Password hashing failed",
		})
		return
	}

	user.Password = string(hash)
	user.ID = len(users) + 1

	// Save user
	users = append(users, user)

	c.JSON(201, gin.H{
		"message": "User registered successfully",
	})
}

func Login(c *gin.Context) {
	var input User

	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request",
		})
		return
	}

	for _, user := range users {
		if user.Email == input.Email {

			err := bcrypt.CompareHashAndPassword(
				[]byte(user.Password),
				[]byte(input.Password),
			)

			if err == nil {
				token := jwt.NewWithClaims(
					jwt.SigningMethodHS256,
					jwt.MapClaims{
						"user_id": user.ID,
						"exp":     time.Now().Add(24 * time.Hour).Unix(),
					},
				)

				tokenString, err := token.SignedString(secretKey)
				if err != nil {
					c.JSON(500, gin.H{
						"error": "Failed to generate token",
					})
					return
				}

				c.JSON(200, gin.H{
					"token": tokenString,
				})
				return
			}
		}
	}

	c.JSON(401, gin.H{
		"error": "Invalid credentials",
	})
}
