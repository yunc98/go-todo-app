package router

import (
	"github.com/yunc98/go-todo-app/controller"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World")
	})

	r.GET("/todo/:id", controller.GetTodo)
	// r.POST("/todo", controller.CreateTodo)

	r.Run()
}