package handlers

import (
	"golang-todo-api/middlewares"
	"golang-todo-api/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetTodos(c *fiber.Ctx) error {
	db := middlewares.GetDB(c)

	var todos []models.Todo
	if err := db.Find(&todos).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve todos")
	}

	return c.JSON(todos)
}

func AddTodo(c *fiber.Ctx) error {
	db := middlewares.GetDB(c)

	// Define a struct to hold the request body data
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	// Generate a new UUID for the ID
	todo.ID = uuid.New().String()

	// Create the new todo in the database
	if err := db.Create(&todo).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create todo")
	}

	return c.Status(fiber.StatusCreated).JSON(todo)
}
