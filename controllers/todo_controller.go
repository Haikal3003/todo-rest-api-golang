package controllers

import (
	"net/http"
	"todo-rest-api/models"

	"github.com/gin-gonic/gin"
)

var todos = []models.Todo{
	{ID: "1", Task: "Belajar Golang", Completed: false},
	{ID: "2", Task: "Mengerjakan tugas", Completed: true},
	{ID: "3", Task: "Membaca buku", Completed: false},
}

func GetTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func GetTodoById(c *gin.Context) {
	id := c.Param("id")
	for _, todo := range todos {
		if todo.ID == id {
			c.IndentedJSON(http.StatusOK, todo)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

func CreateTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Todo berhasil ditambahkan", "data": newTodo})
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var updatedTodo models.Todo

	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos[i] = updatedTodo
			todos[i].ID = id
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Todo berhasil diupdate", "data": updatedTodo})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Todo tidak ditemukan"})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Todo berhasil dihapus"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Todo tidak ditemukan"})
}
