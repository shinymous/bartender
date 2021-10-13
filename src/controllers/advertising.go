package controllers

import (
	"bartender/src/models"
	"bartender/src/repositories"
	"bartender/src/sliceUtils"

	"github.com/gofiber/fiber/v2"
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

	filter := repositories.AdvertisingFilter{
		Resolution:     resolution,
		Format:         format,
		Categorization: params.UserCategorization,
	}

	advertising := repositories.FindAdvertising(filter)
	return c.JSON(advertising.Impression(10.5))
}

func ConfirmImpression(c *fiber.Ctx) error {
	var confirmImpression models.ConfirmImpression
	if err := c.BodyParser(&confirmImpression); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}
	c.Response().SetStatusCode(fiber.StatusNoContent)
	return c.Send(nil)
}
