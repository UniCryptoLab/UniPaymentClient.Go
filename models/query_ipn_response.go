package unipayment

type QueryIPNResponse struct {
	Code string   `json:"code,omitempty"`
	Msg  string   `json:"msg,omitempty"`
	Data []string `json:"data,omitempty"`
}
