package advertising

import (
	"bartender/internal/advertising/models"

	"github.com/gofiber/fiber/v2"
)

type AdvertisingService interface {
	ChooseAdvertising(ctx *fiber.Ctx) error
	ConfirmImpression(ctx *fiber.Ctx) error
	ChooseAdTest(ctx *fiber.Ctx) error
	GenerateUUID(ctx *fiber.Ctx) error
}

type Broker interface {
}

type Topic struct {
	ConfirmImpression string
	RequestAd         string
}

type SendMessage func(topicName string, data interface{})

type BrokerConnection struct {
	SendMessage SendMessage
	Topic       Topic
}

type AdvertisingRepository interface {
	GetUserInfo(ID string) (models.UserInfo, error)
	SaveUserInfo(ID string, userInfo models.UserInfo) error
	FindAdvertising(advertisingFilter models.AdvertisingFilter) models.Advertising
	FindAdvertisingTest() models.Advertising
}
