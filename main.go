package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"todolist/server/handlers"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Static("/", "./client/dist")
	router := app.Group("/api/v1")

	router.Get("/task", handlers.GetAllTask)
	router.Post("/task", handlers.AddTask)
	router.Delete("/task/:id", handlers.DeleteTask)

	app.Listen(":8080")

}