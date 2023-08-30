package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func (ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Hey!",
		})
	})

	app.Listen(":8080")

}