package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/api/v1/expense", func(c *fiber.Ctx) error {
		return c.SendString("Get all expenses")
	})

	app.Post("/api/v1/expense", func(c *fiber.Ctx) error {
		return c.SendString("Create expense")
	})

	app.Listen(":" + os.Getenv("PORT"))
}
