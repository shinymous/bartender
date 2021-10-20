package models

type Advertising struct {
	InternalId     string
	ImpressionId   string
	Name           string
	Creative       string
	Resolution     []uint64
	Format         string
	Categorization string
	Etc            string
	Info           map[string]Criteria
}
type Criteria struct {
	Value  string
	Weight float64
}
type AdvertisingFilter struct {
	Name  string
	Value string
}

func (a Advertising) Impression(delay float64) Impression {
	return Impression{ImpressionId: a.ImpressionId, Creative: a.Creative, Delay: delay, InternalId: a.InternalId}
}
