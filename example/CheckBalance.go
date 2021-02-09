package main

import (
	"log"

	"git.mrcyjanek.net/mrcyjanek/wapi"
)

func main() {
	// First login
	wapi.Login("cyjan", "6290a06de1bc34202f861275b29947a9a9b9fbfcae")
	// Prepare request
	var request wapi.CheckBalanceRequest
	request.Currency = "doge"
	request.Data.Address = "DDbWmzwXqbBSHVdJmwMGaNK1pGUBQDMppY"
	// Send it
	response := wapi.CheckBalance(request)
	if response.OK {
		// Get balance
		log.Println(response.Balance)
	} else {
		// Handle error
		log.Println(response.Message)
	}
}
