package handlers

import (
	"fmt"
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

func GetTodo(c *fiber.Ctx) error {
	db := middlewares.GetDB(c)

	uuid := c.Params("id")

	var todo models.Todo
	if err := db.Where("id = ?", uuid).First(&todo).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Todo not found")
	}

	return c.JSON(todo)
}

func AddTodo(c *fiber.Ctx) error {
	db := middlewares.GetDB(c)

	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	todo.ID = uuid.New().String()

	if err := db.Create(&todo).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create todo")
	}

	return c.Status(fiber.StatusCreated).JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	db := middlewares.GetDB(c)

	uuid := c.Params("id")

	var todo models.Todo
	if err := db.Where("id = ?", uuid).First(&todo).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Todo not found")
	}

	var updatedData models.Todo
	if err := c.BodyParser(&updatedData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	todo.Task = updatedData.Task
	todo.Description = updatedData.Description
	todo.IsFinished = updatedData.IsFinished

	fmt.Println(todo)
	fmt.Println(updatedData)

	if err := db.Save(&todo).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update todo")
	}

	return c.JSON(todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	db := middlewares.GetDB(c)

	uuid := c.Params("id")

	var todo models.Todo
	if err := db.Where("id = ?", uuid).First(&todo).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Todo not found")
	}

	if err := db.Delete(&todo).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete todo")
	}

	return c.Status(fiber.StatusOK).JSON(todo)
}
