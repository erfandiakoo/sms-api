package router

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/")
	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON("It Works!")
	})

	// challenge := api.Group("/challenge")
	// challenge.Post("/login", handler.Login)
	// challenge.Post("/config", handler.Config)

	// ipg := api.Group("/ipg")
	// ipg.Post("/generate", handler.GenerateUrl)
}
