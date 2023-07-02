package controller

import (
	"net/http"
	"strconv"

	"github.com/yunc98/go-todo-app/model"

	"github.com/gin-gonic/gin"
)

func GetTodo(c *gin.Context) {
	i, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter ID"})
		return
	}

	id := uint(i)
	todo := model.Todo{}

	if err := todo.FirstById(id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {
	title := c.PostForm("title")
	completed, err := strconv.ParseBool(c.PostForm("completed"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter 'completed'"})
		return
	}

	todo := model.Todo {
		Title: title,
		Completed: completed,
	}

	todo.Create()

	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	i, err := strconv.Atoi(c.Param("id"))
	id := uint(i)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter ID"})
		return
	}

	todo := model.Todo{}
	if err := todo.FirstById(id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Record not found!"})
		return
	}

	title := c.PostForm("title")
	comp, err := strconv.ParseBool(c.PostForm("completed"))
	completed := bool(comp)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter 'completed'"})
		return
	}

	todo.Title = title
	todo.Completed = completed

	todo.Update()

	c.JSON(http.StatusOK, todo)
}
