package router

import (
	"github.com/erfandiakoo/sms-api/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/")
	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON("It Works!")
	})

	req := api.Group("/api/v1")
	req.Post("/send", handler.SendSMS)
	req.Post("/config", handler.Config)
}
