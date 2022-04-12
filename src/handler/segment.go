package handler

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/matheusroleal/atlas/src/asset"
	"github.com/matheusroleal/atlas/src/cache"
	"github.com/matheusroleal/atlas/src/storage"
	log "github.com/sirupsen/logrus"
)

type NewSegment struct {
	Reference      string
	Identification string
	Data           string
}

func SegmentCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Declare a new Person struct.
	var t NewSegment

	// Try to decode the request body into the struct. If there is an error,
	err := json.NewDecoder(r.Body).Decode(&t)
	// respond to the client with the error message and a 400 status code.
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := asset.CreateAsset(t.Data, t.Identification, t.Reference)

	// Convert asset to a hash
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dataParsed := string(b)
	dataHashed := asset.HashAsset(dataParsed)

	// Send data to relational database
	go storage.InsertSegment("mysql", "root", "password", "Atlas", "tcp(0.0.0.0:6603)", data.ID, dataParsed, data.Reference)

	// Bulk hashed data to send to Blockchain later
	if cache.GetData("0.0.0.0:6379", "", "segments") != "" {
		go cache.AppendData("0.0.0.0:6379", "", "segments", ","+dataHashed)
	} else {
		go cache.SetData("0.0.0.0:6379", "", "segments", dataHashed)
	}

	// Request return
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Segment Created"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Error("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
