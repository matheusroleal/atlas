/*
 * @Author: Matheus Leal
 * @Date: 2022-07-01 22:54:44
 * @Last Modified by: Matheus Leal
 * @Last Modified time: 2022-07-03 14:04:59
 */
package iota

import (
	"math/rand"

	"github.com/davecgh/go-spew/spew"
	. "github.com/iotaledger/iota.go/api"
	log "github.com/sirupsen/logrus"
)

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

/**
 * Get the IOTA node info.
 *
 * @param   endpoint	string		A string with the IOTA address
 * @return  					error			Conection error
 */
func NodeInfo(endpoint string) error {
	// Compose a new API instance
	api, err := ComposeAPI(HTTPClientSettings{URI: endpoint})
	if err != nil {
		log.Error("[IOTA] ERROR: Could not connect to the network")
		return err
	}
	// Get the node info
	nodeInfo, err := api.GetNodeInfo()
	if err != nil {
		log.Error("[IOTA] ERROR: Could not get the node info")
		return err
	}
	// Pretty printer for the response
	spew.Dump(nodeInfo)
	return nil
}

/**
 * Generate a random rune string.
 *
 * @param   n			int				A int with the size of the rune
 * @return  			string		random rune string
 */
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
