package advertising

import (
	"bartender/internal/advertising/models"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/go-redis/redis/v8"
)

type redisClientRepository struct {
	redisClient *redis.Client
}

func NewAdvertisingRepository(redisClient *redis.Client) AdvertisingRepository {
	return &redisClientRepository{
		redisClient: redisClient,
	}
}

func (r *redisClientRepository) GetUserInfo(ID string) (models.UserInfo, error) {
	result, err := r.redisClient.Get(context.Background(), ID).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return models.UserInfo{}, nil
		}
		return models.UserInfo{}, err
	}
	var userInfo models.UserInfo
	if err := json.Unmarshal([]byte(result), &userInfo); err != nil {
		return models.UserInfo{}, err
	}
	return userInfo, nil
}

func (r *redisClientRepository) SaveUserInfo(ID string, userInfo models.UserInfo) error {
	err := r.redisClient.Set(context.Background(), ID, userInfo, 0).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r *redisClientRepository) FindAdvertisingTest() models.AdvertisingTest {
	advertisingTests := []models.AdvertisingTest{
		{
			ImpressionId: "16be3bbe-379d-4917-9f09-998f4d6d58d2",
			CreativeId:   "db477daf-0b6e-4643-8a66-b07042a2dc67",
			PublisherId:  "051ce532-cd18-47a8-ad16-287111cf354d",
			CampaignId:   "12593110-97a7-4f26-bc35-b46554a6602b",
			AdvertiserId: "051ce532-cd18-47a8-ad16-287111cf354d",
			DS:           "g34f4g4sj8923nhd",
		},
		{
			ImpressionId: "16be3bbe-379d-4917-9f09-998f4d6d58d2",
			CreativeId:   "24fa363a-d9e2-47f1-827c-16c04c16d5b8",
			PublisherId:  "051ce532-cd18-47a8-ad16-287111cf354d",
			CampaignId:   "12593110-97a7-4f26-bc35-b46554a6602b",
			AdvertiserId: "051ce532-cd18-47a8-ad16-287111cf354d",
			DS:           "g34f4g4sj8923nhd",
		},
		{
			ImpressionId: "16be3bbe-379d-4917-9f09-998f4d6d58d2",
			CreativeId:   "65d98595-ffcf-44bd-9998-5740da857a4e",
			PublisherId:  "051ce532-cd18-47a8-ad16-287111cf354d",
			CampaignId:   "12593110-97a7-4f26-bc35-b46554a6602b",
			AdvertiserId: "051ce532-cd18-47a8-ad16-287111cf354d",
			DS:           "g34f4g4sj8923nhd",
		},
	}
	return advertisingTests[rand.Intn(2)]
}

func (r *redisClientRepository) FindAdvertising(advertisingFilter []models.AdvertisingFilter) models.Advertising {
	advertisings := []models.Advertising{
		{
			InternalId:   "1",
			ImpressionId: "db477daf-0b6e-4643-8a66-b07042a2dc67",
			Name:         "Ad de Abertura",
			Creative:     "65d98595-ffcf-44bd-9998-5740da857a4e",
			Etc:          "Observações e informações adicionais",
			Info: map[string]models.Criteria{
				"resolutionW": {Value: "1080", Weight: 50},
				"resolutionH": {Value: "720", Weight: 50},
				"format":      {Value: "float", Weight: 1500},
				"keywords":    {Value: []string{"key1", "key2", "key3"}, Weight: 15},
				"gender":      {Value: "M", Weight: 10},
				"yearOfBirth": {Value: "25", Weight: 45},
				"city":        {Value: "São José", Weight: 10},
				"deviceModel": {Value: "Samsung", Weight: 31},
			},
		},
		{
			InternalId:   "2",
			ImpressionId: "db477daf-0b6e-4643-8a66-b07042a2dc67",
			Name:         "Ad de Float",
			Creative:     "89kjkkf21-ffcf-44bd-9998-8891sfffxz2",
			Etc:          "Observações e informações adicionais",
			Info: map[string]models.Criteria{
				"resolutionW": {Value: "3840", Weight: 41},
				"resolutionH": {Value: "2160", Weight: 41},
				"format":      {Value: "tcommerce5", Weight: 1500},
				"keywords":    {Value: []string{"key12", "key21", "key33"}, Weight: 10},
				"gender":      {Value: "F", Weight: 50},
				"yearOfBirth": {Value: "30", Weight: 20},
				"city":        {Value: "Palhoça", Weight: 10},
				"deviceModel": {Value: "TCL", Weight: 15},
			},
		},
		{
			InternalId:   "3",
			ImpressionId: "db4331s-0b6e-4643-8a66-b0s5gffa0",
			Name:         "Ad de Tcommerce5",
			Creative:     "65d98595-ffcf-44bd-9998-5740da857a4e",
			Etc:          "Observações e informações adicionais",
			Info: map[string]models.Criteria{
				"resolutionW": {Value: "1920", Weight: 11},
				"resolutionH": {Value: "1080", Weight: 11},
				"format":      {Value: "opening", Weight: 1500},
				"keywords":    {Value: []string{"key121", "key212", "key333"}, Weight: 70},
				"gender":      {Value: "M", Weight: 100},
				"yearOfBirth": {Value: "55", Weight: 50},
				"city":        {Value: "Florianópolis", Weight: 20},
				"deviceModel": {Value: "LG", Weight: 15},
			},
		},
	}

	var bestScore float64 = 0
	var bestAd models.Advertising = models.Advertising{}

	for _, ad := range advertisings {
		var currentScore float64 = 0
		for _, filter := range advertisingFilter {
			adInfo := ad.Info[filter.Name]
			if isSlice(filter.Value) {
				continue
			}
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

func isSlice(v interface{}) bool {
	switch v.(type) {
	case []string:
		return true
	case string:
		return false
	default:
		return false
	}
}
