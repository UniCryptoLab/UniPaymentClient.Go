package unipayment

type CreateWithdrawalRequest struct {
	Network     string  `json:"network,omitempty"`
	Address     string  `json:"address,omitempty"`
	AssetType   string  `json:"asset_type,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
	DestTag     string  `json:"dest_tag,omitempty"`
	NotifyUrl   string  `json:"notify_url,omitempty"`
	Note        string  `json:"note,omitempty"`
	AutoConfirm bool    `json:"auto_confirm,omitempty"`
	IncludeFee  bool    `json:"include_fee,omitempty"`
}
