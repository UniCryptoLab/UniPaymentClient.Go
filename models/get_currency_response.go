package unipayment

type GetCurrencyResponse struct {
	Code string     `json:"code,omitempty"`
	Msg  string     `json:"msg,omitempty"`
	Data []Currency `json:"data,omitempty"`
}
