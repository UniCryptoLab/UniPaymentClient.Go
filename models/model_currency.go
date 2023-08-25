package unipayment

type Currency struct {
	Code   string `json:"code,omitempty"`
	Symbol string `json:"symbol,omitempty"`
	Name   string `json:"name,omitempty"`
	IsFiat bool   `json:"is_fiat,omitempty"`
}
