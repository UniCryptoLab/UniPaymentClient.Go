package unipayment

type PayoutDetailModel struct {
	PayoutId   string       `json:"payout_id,omitempty"`
	Network    string       `json:"network,omitempty"`
	AssetType  string       `json:"asset_type,omitempty"`
	Status     string       `json:"status,omitempty"`
	CreateTime DateTime     `json:"create_time,omitempty"`
	UpdateTime DateTime     `json:"update_time,omitempty"`
	Items      []PayoutItem `json:"items,omitempty"`
}
