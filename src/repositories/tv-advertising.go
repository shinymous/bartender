package repositories

import (
	"bartender/src/models"
)

type AdvertisingFilter struct {
	Name   string
	Value  string
	Weight float64
}

func FindAdvertising(advertisingFilter []AdvertisingFilter) models.Advertising {
	advertisings := []models.Advertising{
		{
			InternalId:     "1",
			Resolution:     []uint64{1280, 720},
			Format:         "opening",
			Categorization: "Categoria 1",
			ImpressionId:   "db477daf-0b6e-4643-8a66-b07042a2dc67",
			Name:           "Ad de Abertura",
			Creative:       "65d98595-ffcf-44bd-9998-5740da857a4e",
			Etc:            "Observações e informações adicionais",
			Info: map[string]string{
				"resolution":     "1080x720",
				"format":         "format1",
				"categorization": "categorization1",
			},
		},
		{
			InternalId:     "2",
			Resolution:     []uint64{1280, 720},
			Format:         "float",
			Categorization: "Categoria 1",
			ImpressionId:   "db477daf-0b6e-4643-8a66-b07042a2dc67",
			Name:           "Ad de Float",
			Creative:       "89kjkkf21-ffcf-44bd-9998-8891sfffxz2",
			Etc:            "Observações e informações adicionais",
			Info: map[string]string{
				"resolution":     "3840×2160",
				"format":         "format2",
				"categorization": "categorization1",
			},
		},
		{
			InternalId:     "3",
			Resolution:     []uint64{1280, 720},
			Format:         "tcommerce5",
			Categorization: "Categoria 2",
			ImpressionId:   "db4331s-0b6e-4643-8a66-b0s5gffa0",
			Name:           "Ad de Tcommerce5",
			Creative:       "65d98595-ffcf-44bd-9998-5740da857a4e",
			Etc:            "Observações e informações adicionais",
			Info: map[string]string{
				"resolution":     "1920x1080",
				"format":         "format1",
				"categorization": "categorization2",
			},
		},
	}

	var bestScore float64 = 0
	var bestAd models.Advertising = models.Advertising{}

	for _, ad := range advertisings {
		var currentScore float64 = 0
		for _, filter := range advertisingFilter {
			adInfo := ad.Info[filter.Name]
			if adInfo == filter.Value {
				currentScore += filter.Weight
			}
		}
		if currentScore > bestScore {
			bestScore = currentScore
			bestAd = ad
		}
	}
	return bestAd
}
