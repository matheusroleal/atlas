package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/iotaledger/iota.go/trinary"
	"github.com/matheusroleal/atlas/src/asset"
	iotaHandler "github.com/matheusroleal/atlas/src/blockchain/iota"
	storage "github.com/matheusroleal/atlas/src/storage"
)

const endpoint = "https://nodes.devnet.iota.org:443"

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

func main() {
	// Create blocks of data with tags
	for j := 0; j < 31; j++ {
		var tag = RandStringRunes(25)
		log.Println("[Atlas][Test] " + tag)
		// Our data is very long here, it needs to be split over several transactions, 3 in this case
		sum := 0
		var bulk []string
		for i := 0; i < 10; i++ {
			data := asset.CreateAsset(fmt.Sprintf("%s%d", "checkpoint ", sum), fmt.Sprintf("%d", j), "Track1")
			b, err := json.Marshal(data)
			if err != nil {
				panic(err)
			}
			dataParsed := string(b)
			dataHashed := asset.HashAsset(dataParsed)
			// Send data to relational database
			storage.InsertSegment("mysql", "root", "password", "Atlas", "tcp(0.0.0.0:6603)", data.ID, data.Data, data.Reference)
			// Bulk hashed data to be sent to Blockchain
			bulk = append(bulk, dataHashed)
			// log.Println(i)
			sum += i
		}
		log.Println("[Atlas][Test] Sending Messages in a Bulk")
		iotaHandler.BulkData(endpoint, seed, address, bulk, tag)
		t1 := time.Now()
		iotaHandler.RetriveData(endpoint, address, tag)
		t2 := time.Now()
		diff := t2.Sub(t1)
		log.Println(diff)
	}
}
