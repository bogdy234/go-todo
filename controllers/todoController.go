package controllers

import (
	"go/todo-list/initializers"
	"go/todo-list/models"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	// read the body
	var body struct {
		Text      string
		Completed bool
	}
	c.Bind(&body)

	if body.Text == "" {
		c.JSON(400, "You need to provide a text")
		return
	}

	// create a todo with the body values and add it to the database
	todo := models.Todo{
		Text:      body.Text,
		Completed: false,
	}
	result := initializers.DB.Create(&todo)

	// in case of error return 500
	if result.Error != nil {
		c.Status(500)
		return
	}

	// return 200 and result
	c.JSON(200, gin.H{
		"todo": todo,
	})
}

func ReadTodo(c *gin.Context) {
	// read id from params
	id := c.Param("id")

	// find match in the database
	var todo models.Todo
	initializers.DB.First(&todo, id)

	if todo.ID == 0 {
		c.JSON(404, "Todo not found")
	}

	// return 200 code and result
	c.JSON(200, gin.H{
		"todo": todo,
	})
}

func DeleteTodo(c *gin.Context) {
	// read id from params
	id := c.Param("id")

	var todo models.Todo
	initializers.DB.First(&todo, id)

	// guard if TODO is not found
	if todo.ID == 0 {
		c.JSON(404, "Todo not found")
		return
	}

	// delete the entry from db with that id
	result := initializers.DB.Delete(&models.Todo{}, id)

	if result.Error != nil {
		c.JSON(500, result.Error)
	}

	c.JSON(200, "Entry deleted successfully!")
}

func UpdateTodo(c *gin.Context) {
	// get the id from url
	id := c.Param("id")

	// read the body
	var body struct {
		Text      string
		Completed bool
	}
	c.Bind(&body)

	var todo models.Todo
	initializers.DB.First(&todo, id)

	if todo.ID == 0 {
		c.JSON(404, "Todo not found")
		return
	}

	// update the entry with that id
	result := initializers.DB.Model(&todo).Updates(models.Todo{
		Text:      body.Text,
		Completed: body.Completed,
	})

	// check for errors
	if result.Error != nil {
		c.JSON(500, result.Error)
	}

	// return success and updated value
	c.JSON(200, todo)
}
