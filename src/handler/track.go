package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"

	"github.com/iotaledger/iota.go/trinary"
	"github.com/julienschmidt/httprouter"
	"github.com/matheusroleal/atlas/src/asset"
	log "github.com/sirupsen/logrus"

	iotaHandler "github.com/matheusroleal/atlas/src/blockchain/iota"
	"github.com/matheusroleal/atlas/src/storage"
)

// We need a dummy seed even though we don't sign, because the API requires a seed to send
const seed = trinary.Trytes("JBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQ")
const address = trinary.Trytes("XBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQMZOCAHYPD")

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type NewTrack struct {
	Reference      string
	Identification string
}

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
	segments := storage.GetSegment("mysql", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), "Atlas", mysqlEndpoint, t.Reference)
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
	var tag = RandStringRunes(25)
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
