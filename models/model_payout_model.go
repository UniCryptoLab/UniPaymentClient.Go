package unipayment

import (
	"time"
)

type PayoutModel struct {
	PayoutId   string    `json:"payout_id,omitempty"`
	Network    string    `json:"network,omitempty"`
	AssetType  string    `json:"asset_type,omitempty"`
	Status     string    `json:"status,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	UpdateTime time.Time `json:"update_time,omitempty"`
}
