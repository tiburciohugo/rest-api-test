package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiburciohugo/rest-api-test/schemas"
)

func ListTodosHandler(c *gin.Context) {
	todos := []schemas.Todo{}

	if err := db.Find(&todos).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "Error listing todos")
		return
	}
	sendSuccess(c, "list-todos", todos)
}
