package unipayment

type GetPayoutByIdResponse struct {
	Code string             `json:"code,omitempty"`
	Msg  string             `json:"msg,omitempty"`
	Data *PayoutDetailModel `json:"data,omitempty"`
}
