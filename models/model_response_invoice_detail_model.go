package unipayment

type ResponseInvoiceDetailModel struct {
	Code string              `json:"code,omitempty"`
	Msg  string              `json:"msg,omitempty"`
	Data *InvoiceDetailModel `json:"data,omitempty"`
}
