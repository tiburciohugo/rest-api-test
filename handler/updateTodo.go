package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateTodoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Update TodoHandler",
	})

}
