package unipayment

import (
	"time"
)

type InvoiceModel struct {
	Network        string    `json:"network,omitempty"`
	Address        string    `json:"address,omitempty"`
	AppId          string    `json:"app_id,omitempty"`
	InvoiceId      string    `json:"invoice_id,omitempty"`
	OrderId        string    `json:"order_id,omitempty"`
	PriceAmount    float64   `json:"price_amount,omitempty"`
	PriceCurrency  string    `json:"price_currency,omitempty"`
	PayAmount      float64   `json:"pay_amount,omitempty"`
	PayCurrency    string    `json:"pay_currency,omitempty"`
	ExchangeRate   float64   `json:"exchange_rate,omitempty"`
	PaidAmount     float64   `json:"paid_amount,omitempty"`
	CreateTime     time.Time `json:"create_time,omitempty"`
	ExpirationTime time.Time `json:"expiration_time,omitempty"`
	ConfirmSpeed   string    `json:"confirm_speed,omitempty"`
	Status         string    `json:"status,omitempty"`
	ErrorStatus    string    `json:"error_status,omitempty"`
	InvoiceUrl     string    `json:"invoice_url,omitempty"`
}
