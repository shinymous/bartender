package advertising

import "github.com/gofiber/fiber/v2"

type AdvertisingService interface {
	ChooseAdvertising(ctx *fiber.Ctx) error
	ConfirmImpression(ctx *fiber.Ctx) error
}

type Broker interface {
}

type Topic struct {
	ConfirmImpression string
	RequestAd         string
}

type SendAsynMessage func(topicName string, data interface{})

type BrokerConnection struct {
	SendAsynMessage SendAsynMessage
	Topic           Topic
}
