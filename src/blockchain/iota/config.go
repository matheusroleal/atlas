package iota

import (
    . "github.com/iotaledger/iota.go/api"
	"github.com/davecgh/go-spew/spew"
    "log"
)

func NodeInfo(endpoint string) {
	// Compose a new API instance
	api, err := ComposeAPI(HTTPClientSettings{URI: endpoint})
	if err != nil {
		log.Printf("[IOTA] ERROR: Could not connect to the network")
	}
	// Get the node info
	nodeInfo, err := api.GetNodeInfo()
	if err != nil {
		log.Printf("[IOTA] ERROR: Could not get the node info")
	}
	// Pretty printer for the response
	spew.Dump(nodeInfo)
}