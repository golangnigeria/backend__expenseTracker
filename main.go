package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()	
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/api/v1/expense", func(c *fiber.Ctx) error {
		return c.SendString("Get all expenses")
	})

	app.Listen(":3000")
}