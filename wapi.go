package wapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Credentials - Store all the data required for wapi to work.
var Credentials Auth
var wapiurl = "https://mrcyjanek.net/wapi/api/v2.php"

// Login - set the credentials
func Login(username string, APIKey string) {
	Credentials.Username = username
	Credentials.APIKey = APIKey
}

func sendRequest(request interface{}) []byte {
	b, err := json.Marshal(request)
	if err != nil {
		log.Panicln("Failed to Marshal request!", err)
	}
	req, err := http.NewRequest("POST", wapiurl, bytes.NewBuffer(b))
	req.Header.Set("User-Agent", "go wapi (git.mrcyjanek.net/mrcyjanek/wapi)")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
