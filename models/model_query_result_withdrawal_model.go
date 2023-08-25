package unipayment

type QueryResultWithdrawalModel struct {
	Models    []WithdrawalModel `json:"models,omitempty"`
	Total     int32             `json:"total,omitempty"`
	PageNo    int32             `json:"page_no,omitempty"`
	PageCount int32             `json:"page_count,omitempty"`
}
