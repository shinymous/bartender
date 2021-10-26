package models

type Advertising struct {
	ImpressionId string              `json:"impression_id,omitempty"`
	CreativeId   string              `json:"creative_id,omitempty"`
	PublisherId  string              `json:"publisher_id,omitempty"`
	CampaignId   string              `json:"campaign_id,omitempty"`
	AdvertiserId string              `json:"advertiser_id,omitempty"`
	DS           string              `json:"ds,omitempty"`
	Criteria     AdvertisingCriteria `json:"-"`
}

type ReturnAdvertising struct {
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

type YearOfBirth struct {
	Min    int64
	Max    int64
	Weight float64
}

type AdvertisingCriteria struct {
	ResolutionH []Criteria
	ResolutionW []Criteria
	Format      []Criteria
	Keywords    []Criteria
	Gender      []Criteria
	YearOfBirth []YearOfBirth
	City        []Criteria
	DeviceModel []Criteria
}

type AdvertisingFilter struct {
	ResolutionH int64
	ResolutionW int64
	Format      string
	Keywords    []string
	Gender      string
	YearOfBirth int64
	City        string
	DeviceModel string
}
