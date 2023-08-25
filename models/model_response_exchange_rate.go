package unipayment

type ResponseExchangeRate struct {
	Code string        `json:"code,omitempty"`
	Msg  string        `json:"msg,omitempty"`
	Data *ExchangeRate `json:"data,omitempty"`
}
