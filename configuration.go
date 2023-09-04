package unipayment

import (
	"net/http"
)

type AuthParams struct {
	ClientID     string
	ClientSecret string
	AppID        string
}

type Configuration struct {
	BasePath      string            `json:"basePath,omitempty"`
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	HTTPClient    *http.Client
	AuthParams    AuthParams
}

func NewConfiguration(AuthParams AuthParams) *Configuration {
	cfg := &Configuration{
		BasePath:      "https://sandbox-api.unipayment.io",
		DefaultHeader: make(map[string]string),
		UserAgent:     "unipayment_sdk_go/1.1.0",
		AuthParams:    AuthParams,
	}
	return cfg
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}
