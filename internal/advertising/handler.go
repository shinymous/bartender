package advertising

import "github.com/gofiber/fiber/v2"

type AdvertisingHandler struct {
	advertisingService AdvertisingService
}

func NewAdvertisingHandler(advertisingRoute fiber.Router, as AdvertisingService) {
	handler := &AdvertisingHandler{
		advertisingService: as,
	}
	advertisingRoute.Post("/choose-ad", handler.advertisingService.ChooseAdvertising)
	advertisingRoute.Post("/confirm-impression", handler.advertisingService.ConfirmImpression)
	advertisingRoute.Post("/choose-ad-test", handler.advertisingService.ChooseAdTest)
	advertisingRoute.Get("/generate-uuid", handler.advertisingService.GenerateUUID)
}
