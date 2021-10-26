package models

import (
	"encoding/json"
	"time"
)

type UserInfo struct {
	LastImpression string    `json:"LastImpression,omitempty"`
	Timestamp      time.Time `json:"Timestamp,omitempty"`
}

func (i UserInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}
