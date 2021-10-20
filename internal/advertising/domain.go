package advertising

import "github.com/gofiber/fiber/v2"

type AdvertisingService interface {
	ChooseAdvertising(ctx *fiber.Ctx) error
	ConfirmImpression(ctx *fiber.Ctx) error
}

type Broker interface {
}
