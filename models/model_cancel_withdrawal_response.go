package unipayment

type CancelWithdrawalResponse struct {
	Code string       `json:"code,omitempty"`
	Msg  string       `json:"msg,omitempty"`
	Data *interface{} `json:"data,omitempty"`
}
