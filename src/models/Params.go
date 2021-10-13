package models

type Format string

const (
	Float      Format = "float"
	TCommerce5 Format = "tcommerce5"
	Opening    Format = "opening"
)

type Params struct {
	AppId              string                 `json:"app_id,omitempty"`
	UserId             string                 `json:"user_id,omitempty"`
	AllowedFormats     []Format               `json:"allowed_formats,omitempty"`
	UserCategorization string                 `json:"user_categorization,omitempty"`
	UserInfo           map[string]interface{} `json:"user_info,omitempty"`
}

func (i Params) ContainsAllowedFormat(allowedFormat Format) bool {
	for _, a := range i.AllowedFormats {
		if a == allowedFormat {
			return true
		}
	}
	return false
}
