/*
 * @Author: Matheus Leal
 * @Date: 2022-07-01 22:53:33
 * @Last Modified by: Matheus Leal
 * @Last Modified time: 2022-07-03 14:12:55
 */
package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/iotaledger/iota.go/trinary"
	"github.com/julienschmidt/httprouter"
	"github.com/matheusroleal/atlas/src/asset"
	log "github.com/sirupsen/logrus"

	"github.com/matheusroleal/atlas/src/blockchain/iota"
	iotaHandler "github.com/matheusroleal/atlas/src/blockchain/iota"
	"github.com/matheusroleal/atlas/src/storage"
)

// We need a dummy seed even though we don't sign, because the API requires a seed to send
const seed = trinary.Trytes("JBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQ")
const address = trinary.Trytes("XBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQMZOCAHYPD")

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

type NewTrack struct {
	Reference      string
	Identification string
}

/**
 * Track HTTP route handler.
 *
 * @param   w				http.ResponseWriter			The header map that will be sent by WriteHeader
 *					req			http.Request						Specifies the HTTP method
 * @return
 */
func TrackCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Declare a new Person struct.
	var t NewTrack

	// Try to decode the request body into the struct. If there is an error,
	err := json.NewDecoder(r.Body).Decode(&t)
	// respond to the client with the error message and a 400 status code.
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mysqlEndpoint := "tcp(" + os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT") + ")"

	// Get list of Segments related to a Reference ID and compact to one asset
	segments, err := storage.GetSegment("mysql", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), "Atlas", mysqlEndpoint, t.Reference)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	track := asset.CompressAsset(segments)
	data := asset.CreateAsset(track, t.Identification, t.Reference)

	// Convert asset to a hash
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dataParsed := string(b)
	dataHashed := asset.HashAsset(dataParsed)

	// Send data to relational database
	go storage.InsertTrack("mysql", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), "Atlas", mysqlEndpoint, data.ID, dataParsed, data.Reference)

	// Send hashed data to Blockchain
	var tag = iota.RandStringRunes(25)
	go iotaHandler.StoreData(os.Getenv("IOTA_HOST"), seed, address, dataHashed, tag)

	// Request return
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Track Created"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Error("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
