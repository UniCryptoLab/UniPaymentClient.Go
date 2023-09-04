package unipayment

type PayoutRequestItem struct {
	Address string  `json:"address,omitempty"`
	Amount  float64 `json:"amount,omitempty"`
}
