package models

type Advertising struct {
	ImpressionId   string
	Name           string
	Creative       string
	Resolution     []uint64
	Format         string
	Categorization string
	Etc            string
	Info           map[string]string
}

func (a Advertising) Impression(delay float64) Impression {
	return Impression{ImpressionId: a.ImpressionId, Creative: a.Creative, Delay: delay}
}
