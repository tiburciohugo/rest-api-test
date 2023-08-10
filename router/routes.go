package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiburciohugo/rest-api-test/handler"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		v1.GET("/todo", handler.ShowTodoHandler)
		v1.POST("/todo", handler.CreateTodoHandler)
		v1.DELETE("/todo", handler.DeleteTodoHandler)
		v1.PUT("/todo", handler.UpdateTodoHandler)
		v1.GET("/todos", handler.ListTodoHandler)
	}

}
