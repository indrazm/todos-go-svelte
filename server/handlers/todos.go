package handlers

import (
	"github.com/gofiber/fiber/v2"
	"todolist/server/database"
)

func GetAllTask (c *fiber.Ctx) error {
	return c.JSON(database.TodoList)
}