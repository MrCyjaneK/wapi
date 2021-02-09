package wapi

import "encoding/json"

// SendToAddressRequest - pass this to CheckBalance function
type SendToAddressRequest struct {
	APIKey   string `json:"apikey"`   // This one is set by function
	Username string `json:"username"` // This one is set by function
	Action   string `json:"action"`   // This one is set by function
	Currency string `json:"currency"`
	Data     struct {
		Amount  float64 `json:"amount"`
		Address string  `json:"address"`
	} `json:"data"`
}

// SendToAddressResponse - This is what you can get back
type SendToAddressResponse struct {
	OK   bool   `json:"ok"`
	TXID string `json:"txid"`
	// This is empty, unless error
	Code    string `json:"code"`
	Message string `json:"message"`
}

// SendToAddress is used to send money to some address
func SendToAddress(req SendToAddressRequest) SendToAddressResponse {
	req.APIKey = Credentials.APIKey
	req.Username = Credentials.Username
	req.Action = "sendtoaddress"

	b := sendRequest(req)
	var resp SendToAddressResponse
	json.Unmarshal(b, &resp)
	return resp
}
