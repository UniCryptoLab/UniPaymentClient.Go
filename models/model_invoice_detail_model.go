package unipayment

type InvoiceDetailModel struct {
	Transactions []InvoiceTransactionModel `json:"transactions,omitempty"`
}
