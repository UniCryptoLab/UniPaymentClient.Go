package unipayment

type CreateInvoiceResponse struct {
	Code string        `json:"code,omitempty"`
	Msg  string        `json:"msg,omitempty"`
	Data *InvoiceModel `json:"data,omitempty"`
}
