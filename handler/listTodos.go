package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTodosHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET TodoHandler",
	})

}
