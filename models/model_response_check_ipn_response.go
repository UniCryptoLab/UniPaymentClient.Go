package unipayment

type ResponseCheckIpnResponse struct {
	Code string            `json:"code,omitempty"`
	Msg  string            `json:"msg,omitempty"`
	Data *CheckIpnResponse `json:"data,omitempty"`
}
