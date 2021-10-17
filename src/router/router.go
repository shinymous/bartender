package router

import (
	"bartender/src/router/routes"

	"github.com/gofiber/fiber/v2"
)

func Generate() *fiber.App {
	f := fiber.New(fiber.Config{
		CaseSensitive: true,
		ServerHeader:  "Bartender",
		AppName:       "Bartender service v1.0.0-SNAPSHOT",
	})
	return routes.Configure(f)
}
