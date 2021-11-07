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
	broadcaster *string
}

func NewAdvertisingRepository(redisClient *redis.Client, broadcaster *string) AdvertisingRepository {
	return &redisClientRepository{
		redisClient: redisClient,
		broadcaster: broadcaster,
	}
}

func (r *redisClientRepository) GetUserInfo(ID string) (models.UserInfo, error) {
	result, err := r.redisClient.Get(context.Background(), ID+"_"+*r.broadcaster).Result()
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
	err := r.redisClient.Set(context.Background(), ID+"_"+*r.broadcaster, userInfo, 0).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r *redisClientRepository) FindAdvertisingTest() models.Advertising {
	advertisingTests := []models.Advertising{
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

func (r *redisClientRepository) FindAdvertising(advertisingFilter models.AdvertisingFilter) models.Advertising {
	advertisings := []models.Advertising{
		{
			ImpressionId: "16be3bbe-379d-4917-9f09-998f4d6d58d2",
			CreativeId:   "db477daf-0b6e-4643-8a66-b07042a2dc67",
			PublisherId:  "051ce532-cd18-47a8-ad16-287111cf354d",
			CampaignId:   "12593110-97a7-4f26-bc35-b46554a6602b",
			AdvertiserId: "051ce532-cd18-47a8-ad16-287111cf354d",
			DS:           "g34f4g4sj8923nhd",
			Criteria: models.AdvertisingCriteria{
				ResolutionW: []models.Criteria{
					{
						Value: 1920, Weight: 11,
					},
				},
				ResolutionH: []models.Criteria{
					{
						Value: 1080, Weight: 11,
					},
				},
				Format: []models.Criteria{
					{
						Value: "opening", Weight: 1500,
					},
				},
				Keywords: []models.Criteria{
					{
						Value:  "key121",
						Weight: 10,
					},
					{
						Value:  "key212",
						Weight: 20,
					},
					{
						Value:  "key333",
						Weight: 40,
					},
				},
				Gender: []models.Criteria{
					{
						Value:  "M",
						Weight: 100,
					},
					{
						Value:  "F",
						Weight: 20,
					},
				},
				YearOfBirth: []models.YearOfBirth{
					{
						Min:    45,
						Max:    50,
						Weight: 50,
					},
					{
						Min:    10,
						Max:    30,
						Weight: 30,
					},
				},
				City: []models.Criteria{
					{
						Value:  "Florianópolis",
						Weight: 20,
					},
				},
				DeviceModel: []models.Criteria{
					{
						Value:  "LG",
						Weight: 15,
					},
				},
			},
		},
		{
			ImpressionId: "16be3bbe-379d-4917-9f09-998f4d6d58d2",
			CreativeId:   "db477daf-0b6e-4643-8a66-b07042a2dc67",
			PublisherId:  "051ce532-cd18-47a8-ad16-287111cf354d",
			CampaignId:   "12593110-97a7-4f26-bc35-b46554a6602b",
			AdvertiserId: "051ce532-cd18-47a8-ad16-287111cf354d",
			DS:           "g34f4g4sj8923nhd",
			Criteria: models.AdvertisingCriteria{
				ResolutionW: []models.Criteria{
					{
						Value: 1920, Weight: 11,
					},
				},
				ResolutionH: []models.Criteria{
					{
						Value: 1080, Weight: 11,
					},
				},
				Format: []models.Criteria{
					{
						Value: "opening", Weight: 1500,
					},
				},
				Keywords: []models.Criteria{
					{
						Value:  "key121",
						Weight: 10,
					},
					{
						Value:  "key212",
						Weight: 20,
					},
					{
						Value:  "key333",
						Weight: 40,
					},
				},
				Gender: []models.Criteria{
					{
						Value:  "M",
						Weight: 100,
					},
					{
						Value:  "F",
						Weight: 20,
					},
				},
				YearOfBirth: []models.YearOfBirth{
					{
						Min:    45,
						Max:    50,
						Weight: 50,
					},
					{
						Min:    10,
						Max:    30,
						Weight: 30,
					},
				},
				City: []models.Criteria{
					{
						Value:  "Florianópolis",
						Weight: 20,
					},
				},
				DeviceModel: []models.Criteria{
					{
						Value:  "LG",
						Weight: 15,
					},
				},
			},
		},
		{
			ImpressionId: "16be3bbe-379d-4917-9f09-998f4d6d58d2",
			CreativeId:   "db477daf-0b6e-4643-8a66-b07042a2dc67",
			PublisherId:  "051ce532-cd18-47a8-ad16-287111cf354d",
			CampaignId:   "12593110-97a7-4f26-bc35-b46554a6602b",
			AdvertiserId: "051ce532-cd18-47a8-ad16-287111cf354d",
			DS:           "g34f4g4sj8923nhd",
			Criteria: models.AdvertisingCriteria{
				ResolutionW: []models.Criteria{
					{
						Value: 1920, Weight: 11,
					},
				},
				ResolutionH: []models.Criteria{
					{
						Value: 1080, Weight: 11,
					},
				},
				Format: []models.Criteria{
					{
						Value: "opening", Weight: 1500,
					},
				},
				Keywords: []models.Criteria{
					{
						Value:  "key121",
						Weight: 10,
					},
					{
						Value:  "key212",
						Weight: 20,
					},
					{
						Value:  "key333",
						Weight: 40,
					},
				},
				Gender: []models.Criteria{
					{
						Value:  "M",
						Weight: 100,
					},
					{
						Value:  "F",
						Weight: 20,
					},
				},
				YearOfBirth: []models.YearOfBirth{
					{
						Min:    45,
						Max:    50,
						Weight: 50,
					},
					{
						Min:    10,
						Max:    30,
						Weight: 30,
					},
				},
				City: []models.Criteria{
					{
						Value:  "Florianópolis",
						Weight: 20,
					},
				},
				DeviceModel: []models.Criteria{
					{
						Value:  "LG",
						Weight: 15,
					},
				},
			},
		},
	}

	var bestScore float64 = 0
	var bestAd models.Advertising = models.Advertising{}

	for _, ad := range advertisings {
		var currentScore float64 = 0

		for _, city := range ad.Criteria.City {
			if city.Value == advertisingFilter.City {
				currentScore += city.Weight
			}
		}

		for _, deviceModel := range ad.Criteria.DeviceModel {
			if deviceModel.Value == advertisingFilter.DeviceModel {
				currentScore += deviceModel.Weight
			}
		}

		for _, format := range ad.Criteria.Format {
			if format.Value == advertisingFilter.Format {
				currentScore += format.Weight
			}
		}

		for _, gender := range ad.Criteria.Gender {
			if gender.Value == advertisingFilter.Gender {
				currentScore += gender.Weight
			}
		}

		for _, keyword := range ad.Criteria.Keywords {
			for _, filterKeyword := range advertisingFilter.Keywords {
				if keyword.Value == filterKeyword {
					currentScore += keyword.Weight
				}
			}
		}

		for _, resolutionH := range ad.Criteria.ResolutionH {
			if resolutionH.Value == advertisingFilter.ResolutionH {
				currentScore += resolutionH.Weight
			}
		}

		for _, resolutionW := range ad.Criteria.ResolutionW {
			if resolutionW.Value == advertisingFilter.ResolutionW {
				currentScore += resolutionW.Weight
			}
		}

		for _, yearOfBirth := range ad.Criteria.YearOfBirth {
			if advertisingFilter.YearOfBirth > yearOfBirth.Min && advertisingFilter.YearOfBirth < yearOfBirth.Max {
				currentScore += yearOfBirth.Weight
			}
		}
		if currentScore > bestScore {
			bestScore = currentScore
			bestAd = ad
		}
	}

	bestAd.Criteria = models.AdvertisingCriteria{}
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
