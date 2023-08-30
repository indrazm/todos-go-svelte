package handlers

import (
	"strconv"
	"todolist/server/database"
	"todolist/server/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllTask (c *fiber.Ctx) error {
	return c.JSON(database.TodoList)
}

func AddTask (c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	
	database.TodoList = append(database.TodoList, task)
	return c.JSON(database.TodoList)
}

func DeleteTask (c *fiber.Ctx) error {
	id := c.Params("id")

	intValue, _ := strconv.Atoi(id)

	for i, task := range database.TodoList {
		if task.Id == intValue {
			database.TodoList = append(database.TodoList[:i], database.TodoList[i+1:]...)
			break
		}
	}

	return c.JSON(database.TodoList)
}