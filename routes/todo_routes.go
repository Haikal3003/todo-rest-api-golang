package routes

import (
	"todo-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(router *gin.Engine) {
	router.GET("/todos", controllers.GetTodos)
	router.GET("/todos/:id", controllers.GetTodoById)
	router.POST("/todos", controllers.CreateTodo)
	router.PUT("/todos/:id", controllers.UpdateTodo)
	router.DELETE("/todos/:id", controllers.DeleteTodo)
}
