package infrastructure

import (
	"bartender/internal/advertising"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		ServerHeader:  "Bartender",
		AppName:       "Bartender service v1.0.0-SNAPSHOT",
	})
	broker := advertising.CreateConnection()
	advertisingService := advertising.NewAdvertisingService(broker)
	advertising.NewAdvertisingHandler(app.Group("/api"), advertisingService)
	log.Fatal(app.Listen(":8080"))
}
