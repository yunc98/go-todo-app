package router

import (
	"net/http"

	"github.com/yunc98/go-todo-app/controller"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()

	router.Use(CorsMiddleware())

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

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT,DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}