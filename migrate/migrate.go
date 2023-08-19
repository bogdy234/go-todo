package main

import (
	"go/todo-list/initializers"
	"go/todo-list/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()

}

func main() {
	initializers.DB.AutoMigrate(&models.Todo{})
}
