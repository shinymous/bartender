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

// ChooseAdvertising is a function to get best advertising from database
// @Summary Get advertising
// @Description Discover best adivertising
// @Tags advertising
// @Accept json
// @Produce json
// @Param params body models.Params true "bidrequest"
// @Success 200 {object} models.Advertising
// @Failure 500 {object} nil
// @Router /api/v1/choose-ad [post]
func (s *advertisingService) ChooseAdvertising(c *fiber.Ctx) error {
	var params models.Params
	if err := c.BodyParser(&params); err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
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

// ConfirmImpression is a function to notify bartender when a ad was visualized
// @Summary Notify advertising was visualized
// @Description Notify advertising was visualized
// @Tags impression
// @Accept json
// @Produce json
// @Param confirmImpression body models.ConfirmImpression true "ConfirmImpression"
// @Success 204 {object} nil
// @Failure 500 {object} nil
// @Router /api/v1/confirm-impression [post]
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

// ChooseAdTest is a function to get best advertising from database
// @Summary Get advertising test
// @Description Discover best adivertising test
// @Tags advertising test
// @Accept json
// @Produce json
// @Success 200 {object} models.Advertising
// @Failure 500 {object} nil
// @Router /api/v1/choose-ad-test [post]
func (s *advertisingService) ChooseAdTest(c *fiber.Ctx) error {
	return c.JSON(s.repository.FindAdvertisingTest())
}

func SendImpressionInfo(brokerClient BrokerConnection, confirmImpression models.ConfirmImpression) {
	brokerClient.SendAsynMessage(brokerClient.Topic.ConfirmImpression, confirmImpression)
}

func SendRequestInfo(brokerClient BrokerConnection, params models.Params) {
	brokerClient.SendAsynMessage(brokerClient.Topic.RequestAd, params)
}
