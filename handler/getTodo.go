package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowTodoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET TodoHandler",
	})

}
