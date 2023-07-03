package router

import (
	"net/http"

	"github.com/yunc98/go-todo-app/controller"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World")
	})

	router.GET("/todo", controller.GetTodos)
	router.GET("/todo/:id", controller.GetTodoById)
	router.POST("/todo", controller.CreateTodo)
	router.PUT("/todo/:id", controller.UpdateTodoById)
	router.DELETE("/todo/:id", controller.DeleteTodoById)

	router.Run()
}