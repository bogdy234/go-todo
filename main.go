package main

import (
	"go/todo-list/controllers"
	"go/todo-list/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()

}

func main() {

	r := gin.Default()

	r.POST("/todos", controllers.CreateTodo)
	r.GET("/todos/:id", controllers.ReadTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)
	r.PUT("/todos/:id", controllers.UpdateTodo)

	r.Run()
}
