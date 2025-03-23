package main

import (
	"todo-rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.TodoRoutes(router)

	router.Run(":8000")
}
