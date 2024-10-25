package main

import (
	"golang-todo-api/database"
	"golang-todo-api/handlers"
	"golang-todo-api/middlewares"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Run migration and seeding
	database.Init()

	// Initialize Fiber app
	app := fiber.New()

	// Use the dbMiddleware for database connection per request
	app.Use(middlewares.DbMiddleware())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/todos", handlers.GetTodos)
	app.Post("/todo", handlers.AddTodo)

	log.Fatal(app.Listen(":3000"))
}
