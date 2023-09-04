package unipayment

type CheckIpnResponse struct {
	Code string          `json:"code,omitempty"`
	Msg  string          `json:"msg,omitempty"`
	Data *CheckIpnDetail `json:"data,omitempty"`
}
