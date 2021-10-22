package routes

import (
	"bartender/internal/advertising"
	"net/http"
)

func GetAdvertisingRoutes() []Route {
	broker := advertising.CreateConnection()
	advertisingService := advertising.NewAdvertisingService(broker)
	return []Route{
		{
			URI:                   "/choose-ad",
			Method:                http.MethodPost,
			Function:              advertisingService.ChooseAdvertising,
			RequireAuthentication: false,
		},
		{
			URI:                   "/confirm-impression",
			Method:                http.MethodPost,
			Function:              advertisingService.ConfirmImpression,
			RequireAuthentication: false,
		},
	}
}
