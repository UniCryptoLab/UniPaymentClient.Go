package unipayment

type QueryWithdrawalResponse struct {
	Code string                      `json:"code,omitempty"`
	Msg  string                      `json:"msg,omitempty"`
	Data *QueryResultWithdrawalModel `json:"data,omitempty"`
}
