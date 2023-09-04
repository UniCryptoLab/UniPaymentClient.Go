package unipayment

type QueryInvoiceByIdResponse struct {
	Code string              `json:"code,omitempty"`
	Msg  string              `json:"msg,omitempty"`
	Data *InvoiceDetailModel `json:"data,omitempty"`
}
