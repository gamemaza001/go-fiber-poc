package main

import (
	"fmt"

	"reflect"
	"todolist/database"
	"todolist/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func helloWord(c *fiber.Ctx) error {
	return c.SendString("Hello world")
}

func initDatabase() {
	var err error
	dsn := "host=rptcomm.postgres.database.azure.com user=rpt_read password=rpt1234!! dbname=rpt_poc port=5432"
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connected!")
	database.DBConn.AutoMigrate(&models.Todo{})
	fmt.Println("Migrated DB")
}

func setupRoutes(app *fiber.App) {
	app.Get("/todos", models.GetTodos)
	app.Post("/todos", models.CreateTodo)
	app.Get("/books", models.HelloBooks)
}

func main() {
	var c interface{}
	fmt.Println(reflect.ValueOf(&c).Elem().Kind())
	fmt.Println(reflect.ValueOf(&c).Elem().IsZero())
	app := fiber.New()
	initDatabase()
	app.Get("/", helloWord)
	setupRoutes(app)
	var port = os.Getenv("PORT")
	if port == "" {
	   port = ":8080"
	   fmt.Println("No Port In Heroku" + port)
	}
	app.Listen(port)
}
