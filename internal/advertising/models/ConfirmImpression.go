package models

type ConfirmImpression struct {
	ImpressionId      string  `json:"impression_id,omitempty"`
	VisualizationTime float64 `json:"visualization_time,omitempty"`
	Interactions      uint64  `json:"interactions,omitempty"`
}
