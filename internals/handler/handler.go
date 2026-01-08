package handler

import "github.com/gofiber/fiber/v2"

func GetTransactions(ctx *fiber.Ctx) error {
	ctx.JSON(fiber.Map{
		"message": "GetTransactions",
	})
	return nil
}

func PostTransactions(ctx *fiber.Ctx) error {
	ctx.JSON(fiber.Map{
		"message": "PostTransactions",
	})
	return nil
}
