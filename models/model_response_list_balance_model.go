package unipayment

type ResponseListBalanceModel struct {
	Code string         `json:"code,omitempty"`
	Msg  string         `json:"msg,omitempty"`
	Data []BalanceModel `json:"data,omitempty"`
}
