package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golangnigeria/expenseTracker/internals/handler"
)

func Routes(app *fiber.App) *fiber.App {
	app.Get("/v1/transactions", handler.GetTransactions)
	app.Post("/v1/transactions", handler.PostTransactions)

	return app
}