package main

import (
	"log"

	"github.com/erfandiakoo/sms-api/router"
	"github.com/erfandiakoo/sms-api/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	util.InitViper()

	app := fiber.New(fiber.Config{Prefork: false})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,Head,Put,DELETE,PATH",
		AllowHeaders: "",
	}))

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"Erfan": "Diakoo",
			"admin": "@dmin",
		},
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool {
			if user == "Erfan" && pass == "Diakoo" {
				return true
			}
			if user == "admin" && pass == "@dmin" {
				return true
			}
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	}))

	router.SetupRoutes(app)
	err := app.Listen(":8000")
	if err != nil {
		log.Fatalln(err)
	}
}
