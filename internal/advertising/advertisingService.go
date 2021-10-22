package advertising

import (
	"bartender/internal/advertising/models"

	"github.com/gofiber/fiber/v2"
)

type advertisingService struct {
	broker BrokerConnection
}

func NewAdvertisingService(broker BrokerConnection) AdvertisingService {
	return &advertisingService{
		broker: broker,
	}
}

func (s *advertisingService) ChooseAdvertising(c *fiber.Ctx) error {
	return ChooseAdvertising(s.broker, c)
}

func (s *advertisingService) ConfirmImpression(c *fiber.Ctx) error {
	return ConfirmImpression(s.broker, c)
}

func ChooseAdvertising(brokerClient BrokerConnection, c *fiber.Ctx) error {
	var params models.Params
	if err := c.BodyParser(&params); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}
	filter := []models.AdvertisingFilter{
		{
			Name:  "resolution",
			Value: params.UserInfo["resolution"],
		},
		{
			Name:  "format",
			Value: params.UserInfo["format"],
		},
		{
			Name:  "categorization",
			Value: params.UserInfo["categorization"],
		},
	}
	advertising := FindAdvertising(filter)
	go SendRequestInfo(brokerClient, params)
	return c.JSON(advertising.Impression(10.5))
}

func toString(bs []uint64) string {
	b := make([]byte, len(bs))
	for i, v := range bs {
		b[i] = byte(v)
	}
	return string(b)
}

func ConfirmImpression(brokerClient BrokerConnection, c *fiber.Ctx) error {
	var confirmImpression models.ConfirmImpression
	if err := c.BodyParser(&confirmImpression); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	go SendImpressionInfo(brokerClient, confirmImpression)
	c.Response().SetStatusCode(fiber.StatusNoContent)
	return c.Send(nil)
}

func SendImpressionInfo(brokerClient BrokerConnection, confirmImpression models.ConfirmImpression) {
	brokerClient.SendAsynMessage(brokerClient.GetBrokerConfig().CONFIRM_IMPRESSION, confirmImpression)
}

func SendRequestInfo(brokerClient BrokerConnection, params models.Params) {
	brokerClient.SendAsynMessage(brokerClient.GetBrokerConfig().REQUEST_AD, params)
}
