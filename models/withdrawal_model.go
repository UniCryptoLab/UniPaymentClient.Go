package unipayment

import (
	"time"
)

type WithdrawalModel struct {
	Id         string    `json:"id,omitempty"`
	Network    string    `json:"network,omitempty"`
	AssetType  string    `json:"asset_type,omitempty"`
	Amount     float64   `json:"amount,omitempty"`
	Address    string    `json:"address,omitempty"`
	Fee        float64   `json:"fee,omitempty"`
	Status     string    `json:"status,omitempty"`
	TxnHash    string    `json:"txn_hash,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
}
