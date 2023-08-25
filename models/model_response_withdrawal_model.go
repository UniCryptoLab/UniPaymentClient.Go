package unipayment

type ResponseWithdrawalModel struct {
	Code string           `json:"code,omitempty"`
	Msg  string           `json:"msg,omitempty"`
	Data *WithdrawalModel `json:"data,omitempty"`
}
