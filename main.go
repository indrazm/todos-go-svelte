package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"todolist/server/handlers"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", handlers.GetAllTask)
	app.Post("/", handlers.AddTask)
	app.Delete("/:id", handlers.DeleteTask)

	app.Listen(":8080")

}