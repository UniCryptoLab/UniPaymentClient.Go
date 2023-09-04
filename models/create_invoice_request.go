package unipayment

type CreateInvoiceRequest struct {
	Title         string  `json:"title,omitempty"`
	Description   string  `json:"description,omitempty"`
	Lang          string  `json:"lang,omitempty"`
	AppId         string  `json:"app_id,omitempty"`
	PriceAmount   float64 `json:"price_amount,omitempty"`
	PriceCurrency string  `json:"price_currency,omitempty"`
	PayCurrency   string  `json:"pay_currency,omitempty"`
	Network       string  `json:"network,omitempty"`
	NotifyUrl     string  `json:"notify_url,omitempty"`
	RedirectUrl   string  `json:"redirect_url,omitempty"`
	OrderId       string  `json:"order_id,omitempty"`
	ExtArgs       string  `json:"ext_args,omitempty"`
	ConfirmSpeed  string  `json:"confirm_speed,omitempty"`
}
