package unipayment

import (
	"time"
)

type PayoutDetailModel struct {
	PayoutId   string       `json:"payout_id,omitempty"`
	Network    string       `json:"network,omitempty"`
	AssetType  string       `json:"asset_type,omitempty"`
	Status     string       `json:"status,omitempty"`
	CreateTime time.Time    `json:"create_time,omitempty"`
	UpdateTime time.Time    `json:"update_time,omitempty"`
	Items      []PayoutItem `json:"items,omitempty"`
}