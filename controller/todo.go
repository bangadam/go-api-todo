package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyperjiang/gin-skeleton/model"
)

type TodoController struct {}

type CreateTodo struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateTodo struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (ctrl *TodoController) StoreTodo(c *gin.Context) {
	var todo model.Todo
	var todoValidation CreateTodo
	
	if err := c.BindJSON(&todoValidation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.Title = todoValidation.Title
	todo.Description = todoValidation.Description

	if err := todo.Create(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (ctrl *TodoController) UpdateTodo(c *gin.Context) {
	var todo model.Todo
	var todoValidation UpdateTodo

	id := c.Param("id")
	
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
	}

	if err := c.BindJSON(&todoValidation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo.Title = todoValidation.Title
	todo.Description = todoValidation.Description
	
	err := todo.Update(todoValidation.Title, todoValidation.Description, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (ctrl *TodoController) GetTodoById(c *gin.Context) {
	var todo model.Todo
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	err := todo.GetTodoById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (ctrl *TodoController) DeleteTodo(c *gin.Context) {
	var todo model.Todo
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	err := todo.DeleteTodo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "success"})
}