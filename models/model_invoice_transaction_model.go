package unipayment

type InvoiceTransactionModel struct {
	Hash              string  `json:"hash,omitempty"`
	Network           string  `json:"network,omitempty"`
	Symbol            string  `json:"symbol,omitempty"`
	From              string  `json:"from,omitempty"`
	To                string  `json:"to,omitempty"`
	Amount            float64 `json:"amount,omitempty"`
	ConfirmationCount int32   `json:"confirmation_count,omitempty"`
	IsConfirmed       bool    `json:"is_confirmed,omitempty"`
}
