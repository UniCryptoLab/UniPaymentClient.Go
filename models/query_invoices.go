package unipayment

type QueryInvoicesRequest struct {
	AppId     string `schema:"app_id,omitempty"`
	InvoiceId string `schema:"invoice_id,omitempty"`
	OrderId   string `schema:"order_id,omitempty"`
	Status    string `schema:"status,omitempty"`
	PageNo    int    `schema:"page_no"`
	PageSize  int    `schema:"page_size"`
	IsAsc     bool   `schema:"is_asc,omitempty"`
	Start     string `schema:"start,omitempty"`
	End       string `schema:"end,omitempty"`
}
