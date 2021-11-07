package infrastructure

import (
	"bartender/internal/advertising"
	"bartender/internal/infrastructure/broker"
	"bartender/internal/infrastructure/database"
	"fmt"
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func Run() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		ServerHeader:  "Bartender",
		AppName:       "Bartender service v1.0.0-SNAPSHOT",
	})
	broker := broker.CreateConnection()
	redisClient := database.NewRedisClient()
	advertisingRepository := advertising.NewAdvertisingRepository(redisClient)
	advertisingService := advertising.NewAdvertisingService(
		advertising.BrokerConnection{SendAsynMessage: broker.SendAsynMessage,
			Topic: advertising.Topic{
				ConfirmImpression: broker.GetBrokerConfig().CONFIRM_IMPRESSION,
				RequestAd:         broker.GetBrokerConfig().REQUEST_AD,
			}},
		advertisingRepository,
	)
	// app.Group("/api", middleware)
	advertising.NewAdvertisingHandler(app.Group("/api/v1"), advertisingService)
	app.Get("/docs/*", swagger.Handler)
	log.Fatal(app.Listen(":8080"))
}

func middleware(c *fiber.Ctx) error {
	fmt.Println("Alterar parametros de acordo com a emissora!")
	c.Next()
	return nil
}
