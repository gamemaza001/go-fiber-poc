package models

import (
	"todolist/database"
	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Book struct {
	Unique_id  string 
	Name  string 
	Age int 
}

var books = []Book{
	{Unique_id: "1", Name: "Harry Potter", Age: 24},
	{Unique_id: "2", Name: "The Lord of the Rings", Age: 34},
	{Unique_id: "3", Name: "The Wizard of Oz", Age: 15},
}

func GetTodos(c *fiber.Ctx) error {
	db := database.DBConn
	var todos []Todo
	db.Find(&todos)
	return c.JSON(&todos)
}

func CreateTodo(c *fiber.Ctx) error {
	db := database.DBConn
	todo := new(Todo)
	err := c.BodyParser(todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "massage": "Check your input", "data": err})
	}
	err = db.Create(&todo).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "massage": "Could not create todo", "data": err})
	}
	return c.JSON(&todo)
}

func HelloBooks(c *fiber.Ctx) error {
	return c.JSON(&books)
}

