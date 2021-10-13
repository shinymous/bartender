package channels

import (
	"bartender/src/models"
)

var AdvertisingChannel chan models.Advertising

func InitAdvertisingChannel() {
	AdvertisingChannel = make(chan models.Advertising)
}
