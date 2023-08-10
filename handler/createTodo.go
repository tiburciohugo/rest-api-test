package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTodoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "POST TodoHandler",
	})

}
