package unipayment

type GetExchangeRateByFiatCurrencyResponse struct {
	Code string         `json:"code,omitempty"`
	Msg  string         `json:"msg,omitempty"`
	Data []ExchangeRate `json:"data,omitempty"`
}
