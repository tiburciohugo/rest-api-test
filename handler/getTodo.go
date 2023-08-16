package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiburciohugo/rest-api-test/schemas"
)

func ShowTodoHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	todo := schemas.Todo{}

	if err := db.First(&todo, id).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("todo with id: %s not found", id))
		return
	}
	sendSuccess(c, "todo found successfully", todo)
}
