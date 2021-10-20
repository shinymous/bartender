package advertising

import "bartender/internal/advertising/models"

func FindAdvertising(advertisingFilter []models.AdvertisingFilter) models.Advertising {
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
			Info: map[string]models.Criteria{
				"resolution":     {Value: "1080x720", Weight: 50},
				"format":         {Value: "format1", Weight: 10},
				"categorization": {Value: "categorization1", Weight: 10},
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
			Info: map[string]models.Criteria{
				"resolution":     {Value: "3840×2160", Weight: 50},
				"format":         {Value: "format2", Weight: 10},
				"categorization": {Value: "categorization1", Weight: 10},
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
			Info: map[string]models.Criteria{
				"resolution":     {Value: "1920x1080", Weight: 50},
				"format":         {Value: "format1", Weight: 10},
				"categorization": {Value: "categorization2", Weight: 10},
			},
		},
	}

	var bestScore float64 = 0
	var bestAd models.Advertising = models.Advertising{}

	for _, ad := range advertisings {
		var currentScore float64 = 0
		for _, filter := range advertisingFilter {
			adInfo := ad.Info[filter.Name]
			if adInfo.Value == filter.Value {
				currentScore += adInfo.Weight
			}
		}
		if currentScore > bestScore {
			bestScore = currentScore
			bestAd = ad
		}
	}
	return bestAd
}
