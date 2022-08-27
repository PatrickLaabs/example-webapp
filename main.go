package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/gin", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "GIN GIN!"})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	r.Run("0.0.0.0:9090")
}
