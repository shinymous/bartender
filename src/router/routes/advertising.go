package routes

import (
	"bartender/src/controllers"
	"net/http"
)

var advertisingRoutes = []Route{
	{
		URI:                   "/choose-ad",
		Method:                http.MethodPost,
		Function:              controllers.ChooseAdvertising,
		RequireAuthentication: false,
	},
	{
		URI:                   "/confirm-impression",
		Method:                http.MethodPost,
		Function:              controllers.ConfirmImpression,
		RequireAuthentication: false,
	},
}
