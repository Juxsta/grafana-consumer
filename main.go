package main

import (
	"net/http"

	"github.com/Juxsta/grafana-consumer/api" // Corrected import path

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Apply the RequestLogger middleware
	r.Use(api.RequestLogger())

	// Register API routes
	api.SetupRoutes(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}
