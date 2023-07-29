package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()

	// Define a sample route
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})

	// Run the server
	r.Run(":8080")
}
