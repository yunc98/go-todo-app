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

	router.GET("/todo/:id", controller.GetTodo)
	router.POST("/todo", controller.CreateTodo)

	router.Run()
}