package iota

import (
	"github.com/davecgh/go-spew/spew"
	. "github.com/iotaledger/iota.go/api"
	log "github.com/sirupsen/logrus"
)

func NodeInfo(endpoint string) {
	// Compose a new API instance
	api, err := ComposeAPI(HTTPClientSettings{URI: endpoint})
	if err != nil {
		log.Error("[IOTA] ERROR: Could not connect to the network")
	}
	// Get the node info
	nodeInfo, err := api.GetNodeInfo()
	if err != nil {
		log.Error("[IOTA] ERROR: Could not get the node info")
	}
	// Pretty printer for the response
	spew.Dump(nodeInfo)
}
