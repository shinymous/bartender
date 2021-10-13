package repositories

import (
	"bartender/src/models"
	"math/rand"
)

type AdvertisingFilter struct {
	Resolution     []uint64
	Format         string
	Categorization string
}

func FindAdvertising(advertisingFilter AdvertisingFilter) models.Advertising {
	advertisings := []models.Advertising{
		{
			Resolution:     []uint64{1280, 720},
			Format:         "opening",
			Categorization: "Categoria 1",
			ImpressionId:   "db477daf-0b6e-4643-8a66-b07042a2dc67",
			Name:           "Ad de Abertura",
			Creative:       "65d98595-ffcf-44bd-9998-5740da857a4e",
			Etc:            "Observações e informações adicionais",
		},
		{
			Resolution:     []uint64{1280, 720},
			Format:         "float",
			Categorization: "Categoria 1",
			ImpressionId:   "db477daf-0b6e-4643-8a66-b07042a2dc67",
			Name:           "Ad de Float",
			Creative:       "89kjkkf21-ffcf-44bd-9998-8891sfffxz2",
			Etc:            "Observações e informações adicionais",
		},
		{
			Resolution:     []uint64{1280, 720},
			Format:         "tcommerce5",
			Categorization: "Categoria 2",
			ImpressionId:   "db4331s-0b6e-4643-8a66-b0s5gffa0",
			Name:           "Ad de Tcommerce5",
			Creative:       "65d98595-ffcf-44bd-9998-5740da857a4e",
			Etc:            "Observações e informações adicionais",
		},
	}
	return advertisings[rand.Intn(2-0)+0]
}
