package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiburciohugo/rest-api-test/schemas"
)

// @BasePath /api/v1

// @Summary Create a todo
// @Description Create a new todo
// @Tags Todos
// @Accept  json
// @Produce  json
// @Param request body CreateTodoRequest true "Request Body"
// @Success 200 {object} CreateTodoResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /todo [post]

func CreateTodoHandler(c *gin.Context) {
	request := CreateTodoRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Error validating request: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	todo := schemas.Todo{
		Title:       request.Title,
		Description: request.Description,
		Completed:   request.Completed,
	}

	if err := db.Create(&request).Error; err != nil {
		logger.Errorf("Error creating todo: %v", err.Error())
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, "create-todo", todo)
}
