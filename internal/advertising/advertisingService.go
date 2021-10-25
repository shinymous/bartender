package advertising

import (
	"bartender/internal/advertising/models"
	"strconv"
	"strings"

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
			Name:  "resolutionW",
			Value: strconv.Itoa(int(params.Context.Device.W)),
		},
		{
			Name:  "resolutionH",
			Value: strconv.Itoa(int(params.Context.Device.H)),
		},
		{
			Name:  "format",
			Value: params.Item[0].Spec.TagID,
		},
		{
			Name:  "keywords",
			Value: strings.Split(params.Context.User.Keywords, ","),
		},
		{
			Name:  "gender",
			Value: params.Context.User.Gender,
		},
		{
			Name:  "yearOfBirth",
			Value: strconv.Itoa(int(params.Context.User.Yob)),
		},
		{
			Name:  "city",
			Value: params.Context.User.Geo.City,
		},
		{
			Name:  "deviceModel",
			Value: params.Context.Device.Model,
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
