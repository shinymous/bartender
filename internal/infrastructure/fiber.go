package infrastructure

import (
	"bartender/internal/advertising"
	"bartender/internal/infrastructure/broker"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		ServerHeader:  "Bartender",
		AppName:       "Bartender service v1.0.0-SNAPSHOT",
	})
	broker := broker.CreateConnection()
	advertisingService := advertising.NewAdvertisingService(advertising.BrokerConnection{SendAsynMessage: broker.SendAsynMessage,
		Topic: advertising.Topic{
			ConfirmImpression: broker.GetBrokerConfig().CONFIRM_IMPRESSION,
			RequestAd:         broker.GetBrokerConfig().REQUEST_AD,
		}},
	)
	app.Group("/", middleware)
	advertising.NewAdvertisingHandler(app.Group("/api"), advertisingService)
	log.Fatal(app.Listen(":8080"))
}

func middleware(c *fiber.Ctx) error {
	fmt.Println("Alterar parametros de acordo com a emissora!")
	c.Next()
	return nil
}
