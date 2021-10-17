package controllers

import (
	"bartender/src/db"
	"bartender/src/models"
	"bartender/src/repositories"

	"github.com/gofiber/fiber/v2"
)

func ChooseAdvertising(c *fiber.Ctx) error {
	var params models.Params
	if err := c.BodyParser(&params); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}
	filter := []repositories.AdvertisingFilter{
		{
			Name:   "resolution",
			Value:  params.UserInfo["resolution"],
			Weight: 20,
		},
		{
			Name:   "format",
			Value:  params.UserInfo["format"],
			Weight: 10,
		},
		{
			Name:   "categorization",
			Value:  params.UserInfo["categorization"],
			Weight: 30,
		},
	}
	advertising := repositories.FindAdvertising(filter)
	go SendRequestInfo(params)
	return c.JSON(advertising.Impression(10.5))
}

func toString(bs []uint64) string {
	b := make([]byte, len(bs))
	for i, v := range bs {
		b[i] = byte(v)
	}
	return string(b)
}

func ConfirmImpression(c *fiber.Ctx) error {
	var confirmImpression models.ConfirmImpression
	if err := c.BodyParser(&confirmImpression); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	go SendImpressionInfo(confirmImpression)
	c.Response().SetStatusCode(fiber.StatusNoContent)
	return c.Send(nil)
}

func SendImpressionInfo(confirmImpression models.ConfirmImpression) {
	db.SendAsynMessage(db.CONFIRM_IMPRESSION, confirmImpression)
}

func SendRequestInfo(params models.Params) {
	db.SendAsynMessage(db.REQUEST_AD, params)
}
