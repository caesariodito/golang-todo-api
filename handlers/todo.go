package handlers

import (
	"golang-todo-api/middlewares"
	"golang-todo-api/models"

	"github.com/gofiber/fiber/v2"
)

func GetTodos(c *fiber.Ctx) error {
	db := middlewares.GetDB(c)

	var todos []models.Todo
	if err := db.Find(&todos).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve todos")
	}

	return c.JSON(todos)
}
