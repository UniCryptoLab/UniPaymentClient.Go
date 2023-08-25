package unipayment

type BalanceModel struct {
	AssetType     string  `json:"asset_type,omitempty"`
	Balance       float64 `json:"balance,omitempty"`
	FrozenBalance float64 `json:"frozen_balance,omitempty"`
	Available     float64 `json:"available,omitempty"`
}
