package unipayment

type ResponseQueryResultPayoutModel struct {
	Code string                  `json:"code,omitempty"`
	Msg  string                  `json:"msg,omitempty"`
	Data *QueryResultPayoutModel `json:"data,omitempty"`
}
