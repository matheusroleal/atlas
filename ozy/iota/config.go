package iota

import (
    "github.com/iotaledger/iota.go/api"
    "log"
)

func Connect(address string) {
	// Compose a new API instance
	api, err := ComposeAPI(HTTPClientSettings{URI: address})
    if err != nil {
		log.Printf("[IOTA] ERROR: Could not connect to the network")
		return nil
    }
	return api
}