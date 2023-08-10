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
func ShowTodoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET TodoHandler",
	})

}
func DeleteTodoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete TodoHandler",
	})

}
func UpdateTodoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Update TodoHandler",
	})

}
func ListTodoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET TodoHandler",
	})

}
