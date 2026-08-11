package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Authentication API",
			"Version": "1.0.0",
		})
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server")

	}
}
