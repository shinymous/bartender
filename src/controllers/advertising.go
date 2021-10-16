package controllers

import (
	"bartender/src/db"
	"bartender/src/models"
	"bartender/src/repositories"
	"bartender/src/sliceUtils"
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"github.com/jeroenrinzema/commander"
)

func ChooseAdvertising(c *fiber.Ctx) error {
	var params models.Params
	if err := c.BodyParser(&params); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}
	var resolution []uint64
	if params.UserInfo != nil && params.UserInfo["resolution"] != nil {
		genericResolutions := sliceUtils.InterfaceSlice(params.UserInfo["resolution"])
		resolution := make([]uint64, len(genericResolutions))
		for i, value := range genericResolutions {
			val, ok := value.(uint64)
			if ok {
				resolution[i] = val
			}
		}
	}
	var format string
	allowedFormatLen := len(params.AllowedFormats)
	if allowedFormatLen > 0 {
		if params.ContainsAllowedFormat(models.Opening) {
			format = string(models.Opening)
		} else {
			if len(params.AllowedFormats) == 1 {
				format = string(params.AllowedFormats[0])
			} else {
				format = string(models.Float)
			}
		}
	}
	filter := []repositories.AdvertisingFilter{
		{
			Name:   "Resolution",
			Value:  toString(resolution),
			Weight: 20,
		},
		{
			Name:   "Format",
			Value:  format,
			Weight: 10,
		},
		{
			Name:   "Categorization",
			Value:  params.UserCategorization,
			Weight: 30,
		},
	}

	advertising := repositories.FindAdvertising(filter)
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
	key, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
	}
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(confirmImpression)
	command := commander.NewMessage("", 0, key.Bytes(), reqBodyBytes.Bytes())
	err = db.ConfirmImpressionMessageBrokerClient.AsyncCommand(command)
	if err != nil {
		fmt.Println(err)
	}
}
