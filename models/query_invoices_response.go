package unipayment

type QueryInvoicesResponse struct {
	Code string                   `json:"code,omitempty"`
	Msg  string                   `json:"msg,omitempty"`
	Data *QueryResultInvoiceModel `json:"data,omitempty"`
}
