package main

import (
	"log"

    "github.com/iotaledger/iota.go/api"
    "github.com/davecgh/go-spew/spew"

	iota "github.com/matheusroleal/ozy/blockchain/iota"
)

var endpoint = "https://nodes.devnet.thetangle.org"

func main() {
	api = iota.Connect(endpoint)

    nodeInfo, err := api.GetNodeInfo()
    if err != nil {
		log.Printf("[IOTA] ERROR: Could not get the node info")
    }

    spew.Dump(nodeInfo)
}