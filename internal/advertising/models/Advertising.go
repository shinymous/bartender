package models

type Advertising struct {
	InternalId   string
	ImpressionId string
	Name         string
	Creative     string
	Etc          string
	Info         map[string]Criteria
}

type AdvertisingTest struct {
	ImpressionId string `json:"impression_id,omitempty"`
	CreativeId   string `json:"creative_id,omitempty"`
	PublisherId  string `json:"publisher_id,omitempty"`
	CampaignId   string `json:"campaign_id,omitempty"`
	AdvertiserId string `json:"advertiser_id,omitempty"`
	DS           string `json:"ds,omitempty"`
}

type Criteria struct {
	Value  interface{}
	Weight float64
}
type AdvertisingFilter struct {
	Name  string
	Value interface{}
}

func (a Advertising) Impression(delay float64) Impression {
	return Impression{ImpressionId: a.ImpressionId, Creative: a.Creative, Delay: delay, InternalId: a.InternalId}
}
