package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tiburciohugo/rest-api-test/docs"
	"github.com/tiburciohugo/rest-api-test/handler"
)

func initializeRoutes(r *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := r.Group(basePath)
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
		v1.GET("/todos", handler.ListTodosHandler)
	}
	// Initialize Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
