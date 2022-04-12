package handler

import (
	"encoding/json"
	"net/http"
	"os"

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

	mysqlEndpoint := "tcp(" + os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT") + ")"
	// Send data to relational database
	go storage.InsertSegment("mysql", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), "Atlas", mysqlEndpoint, data.ID, dataParsed, data.Reference)

	redisEndpoint := os.Getenv("REDIS_HOST") + ":" + os.Getenv("PASSWORD_HOST")
	// Bulk hashed data to send to Blockchain later
	if cache.GetData(redisEndpoint, "", "segments") != "" {
		go cache.AppendData(redisEndpoint, "", "segments", ","+dataHashed)
	} else {
		go cache.SetData(redisEndpoint, "", "segments", dataHashed)
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
