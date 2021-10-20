package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	URI                   string
	Method                string
	Function              func(*fiber.Ctx) error
	RequireAuthentication bool
}

func Configure(f *fiber.App) *fiber.App {
	routes := GetAdvertisingRoutes()
	api := f.Group("/", middleware)
	for _, route := range routes {
		api.Add(route.Method, route.URI, route.Function)
	}
	return f
}

func middleware(c *fiber.Ctx) error {
	fmt.Println("Alterar parametros de acordo com a emissora!")
	c.Next()
	return nil
}
