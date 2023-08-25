package unipayment

type CreatePayoutRequest struct {
	Items     []PayoutRequestItem `json:"items,omitempty"`
	Network   string              `json:"network,omitempty"`
	AssetType string              `json:"asset_type,omitempty"`
}
