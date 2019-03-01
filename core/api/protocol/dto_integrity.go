package protocol

type IntegrityResponse struct {
	Message   string `json:"message"`
	Millis    string `json:"millis"`
	Hash      string `json:"hash"`
	Signature string `json:"signature"`
}
