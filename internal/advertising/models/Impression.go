package models

type Impression struct {
	ImpressionId string  `json:"impression_id,omitempty"`
	Creative     string  `json:"creative,omitempty"`
	Delay        float64 `json:"delay,omitempty"`
	InternalId   string  `json:"internal_id,omitempty"`
}
