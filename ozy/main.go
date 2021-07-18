package main

import (
	"log"

  "github.com/iotaledger/iota.go/api"
  "github.com/davecgh/go-spew/spew"

	iotaHandler "github.com/matheusroleal/ozymandias/ozy/iota"
)

var endpoint = "https://nodes.devnet.thetangle.org"

func main() {
	api = iotaHandler.Connect(endpoint)

  nodeInfo, err := api.GetNodeInfo()
  if err != nil {
  log.Printf("[IOTA] ERROR: Could not get the node info")
  }

  spew.Dump(nodeInfo)
}