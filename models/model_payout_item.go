package unipayment

type PayoutItem struct {
	Address string  `json:"address,omitempty"`
	Amount  float64 `json:"amount,omitempty"`
	Hash    string  `json:"hash,omitempty"`
}
