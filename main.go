package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	InitDB()

	r := gin.Default()

	// Home Route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Authentication APIs
	r.POST("/auth/register", Register)
	r.POST("/auth/login", Login)

	// Protected Routes
	auth := r.Group("/")
	auth.Use(AuthMiddleware())

	auth.POST("/tickets", CreateTicket)
	auth.GET("/tickets", GetTickets)
	auth.GET("/tickets/:id", GetTicket)
	auth.PATCH("/tickets/:id/status", UpdateStatus)

	// Render PORT support
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
