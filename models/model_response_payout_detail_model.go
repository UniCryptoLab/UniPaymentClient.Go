package unipayment

type ResponsePayoutDetailModel struct {
	Code string             `json:"code,omitempty"`
	Msg  string             `json:"msg,omitempty"`
	Data *PayoutDetailModel `json:"data,omitempty"`
}
