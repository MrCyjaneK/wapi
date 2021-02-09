package wapi

import "encoding/json"

// CheckBalanceRequest - pass this to CheckBalance function
type CheckBalanceRequest struct {
	APIKey   string `json:"apikey"`   // This one is set by function
	Username string `json:"username"` // This one is set by function
	Action   string `json:"action"`   // This one is set by function
	Currency string `json:"currency"`
	Data     struct {
		Address string `json:"address"`
	} `json:"data"`
}

// CheckBalanceResponse - This is what you can get back
type CheckBalanceResponse struct {
	OK       bool    `json:"ok"`
	Balance  float64 `json:"balance"`
	Received float64 `json:"received"`
	// This is empty, unless error
	Code    string `json:"code"`
	Message string `json:"message"`
}

// CheckBalance is used to check balance of one address
func CheckBalance(req CheckBalanceRequest) CheckBalanceResponse {
	req.APIKey = Credentials.APIKey
	req.Username = Credentials.Username
	req.Action = "checkbalance"

	b := sendRequest(req)
	var resp CheckBalanceResponse
	json.Unmarshal(b, &resp)
	return resp
}
