package unipayment

type GetWalletBalancesResponse struct {
	Code string         `json:"code,omitempty"`
	Msg  string         `json:"msg,omitempty"`
	Data []BalanceModel `json:"data,omitempty"`
}
