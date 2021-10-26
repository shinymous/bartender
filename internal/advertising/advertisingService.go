package advertising

import (
	"bartender/internal/advertising/models"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type advertisingService struct {
	broker     BrokerConnection
	repository AdvertisingRepository
}

func NewAdvertisingService(broker BrokerConnection, repository AdvertisingRepository) AdvertisingService {
	return &advertisingService{
		broker:     broker,
		repository: repository,
	}
}

func (s *advertisingService) ChooseAdvertising(c *fiber.Ctx) error {
	var params models.Params
	if err := c.BodyParser(&params); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	userInfo, err := s.repository.GetUserInfo(params.Context.User.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if (userInfo == models.UserInfo{} || userInfo != models.UserInfo{}) {
		diff := time.Now().Sub(userInfo.Timestamp)
		if diff.Minutes() < 10 {
			return nil
		}
	}

	filter := models.AdvertisingFilter{
		ResolutionH: params.Context.Device.H,
		ResolutionW: params.Context.Device.W,
		Format:      params.Item[0].Spec.TagID,
		Keywords:    strings.Split(params.Context.User.Keywords, ","),
		Gender:      params.Context.User.Gender,
		YearOfBirth: params.Context.User.Yob,
		City:        params.Context.User.Geo.City,
		DeviceModel: params.Context.Device.Model,
	}

	advertising := s.repository.FindAdvertising(filter)
	go SendRequestInfo(s.broker, params)
	go s.repository.SaveUserInfo(params.Context.User.ID, models.UserInfo{Timestamp: time.Now(), LastImpression: advertising.ImpressionId})
	return c.JSON(advertising)
}

func (s *advertisingService) ConfirmImpression(c *fiber.Ctx) error {
	var confirmImpression models.ConfirmImpression
	if err := c.BodyParser(&confirmImpression); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	go SendImpressionInfo(s.broker, confirmImpression)
	c.Response().SetStatusCode(fiber.StatusNoContent)
	return c.Send(nil)
}

func toString(bs []uint64) string {
	b := make([]byte, len(bs))
	for i, v := range bs {
		b[i] = byte(v)
	}
	return string(b)
}

func (s *advertisingService) ChooseAdTest(c *fiber.Ctx) error {
	return c.JSON(s.repository.FindAdvertisingTest())
}

func SendImpressionInfo(brokerClient BrokerConnection, confirmImpression models.ConfirmImpression) {
	brokerClient.SendAsynMessage(brokerClient.Topic.ConfirmImpression, confirmImpression)
}

func SendRequestInfo(brokerClient BrokerConnection, params models.Params) {
	brokerClient.SendAsynMessage(brokerClient.Topic.RequestAd, params)
}
