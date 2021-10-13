package routes

import (
	"github.com/gofiber/fiber/v2"
)

type Route struct {
	URI                   string
	Method                string
	Function              func(*fiber.Ctx) error
	RequireAuthentication bool
}

func Configure(f *fiber.App) *fiber.App {
	routes := advertisingRoutes
	for _, route := range routes {
		f.Add(route.Method, route.URI, route.Function)
	}
	return f
}
