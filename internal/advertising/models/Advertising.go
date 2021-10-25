package models

type Advertising struct {
	InternalId   string
	ImpressionId string
	Name         string
	Creative     string
	Etc          string
	Info         map[string]Criteria
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
