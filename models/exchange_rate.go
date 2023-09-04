package unipayment

type ExchangeRate struct {
	From string  `json:"from,omitempty"`
	To   string  `json:"to,omitempty"`
	Rate float64 `json:"rate,omitempty"`
}
