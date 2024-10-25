package middlewares

import (
	"context"
	"golang-todo-api/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DbMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		db, err := database.ConnectDB()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Database connection failed")
		}
		defer database.DisconnectDB(db)

		// Pass the db connection through context
		ctx := context.WithValue(c.UserContext(), "db", db)
		c.SetUserContext(ctx)

		return c.Next()
	}
}

func GetDB(c *fiber.Ctx) *gorm.DB {
	return c.UserContext().Value("db").(*gorm.DB)
}
