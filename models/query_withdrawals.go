package unipayment

type QueryWithdrawalRequest struct {
	Network   string `schema:"network,omitempty"`
	AssetType string `schema:"asset_type,omitempty"`
	Status    string `schema:"status,omitempty"`
	PageNo    int    `schema:"page_no"`
	PageSize  int    `schema:"page_size"`
	IsAsc     bool   `schema:"is_asc,omitempty"`
	Start     string `schema:"start,omitempty"`
	End       string `schema:"end,omitempty"`
}
