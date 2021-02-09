package wapi

import "encoding/json"

// GetNewAddressRequest - pass this to CheckBalance function
type GetNewAddressRequest struct {
	APIKey   string   `json:"apikey"`   // This one is set by function
	Username string   `json:"username"` // This one is set by function
	Action   string   `json:"action"`   // This one is set by function
	Currency string   `json:"currency"`
	Data     struct{} `json:"data"`
}

// GetNewAddressResponse - This is what you can get back
type GetNewAddressResponse struct {
	OK      bool   `json:"ok"`
	Address string `json:"address"`
	// This is empty, unless error
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GetNewAddress is used to generate new deposit address
func GetNewAddress(req GetNewAddressRequest) GetNewAddressResponse {
	req.APIKey = Credentials.APIKey
	req.Username = Credentials.Username
	req.Action = "getnewaddress"

	b := sendRequest(req)
	var resp GetNewAddressResponse
	json.Unmarshal(b, &resp)
	return resp
}
