package test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/matheusroleal/atlas/src/asset"
	iotaHandler "github.com/matheusroleal/atlas/src/blockchain/iota"
	log "github.com/sirupsen/logrus"
)

func searchInsertionTest() {
	// Create blocks of data with tags
	for j := 0; j < 31; j++ {
		var tag = RandStringRunes(25)
		log.Debug("[Atlas][Test] " + tag)
		// Our data is very long here, it needs to be split over several transactions, 3 in this case
		sum := 0
		var segments []asset.Asset
		for i := 0; i < 100; i++ {
			data := asset.CreateAsset(fmt.Sprintf("%s%d", "checkpoint ", sum), fmt.Sprintf("%d", j), "Track1")
			// Bulk hashed data to be sent to Blockchain
			segments = append(segments, *data)
			sum += i
		}

		track := asset.CompressAsset(segments)
		result := asset.CreateAsset(track, fmt.Sprint("%s%d", "id", j), "track1")

		// Convert asset to a hash
		b, err := json.Marshal(result)
		if err != nil {
			return
		}
		dataParsed := string(b)
		dataHashed := asset.HashAsset(dataParsed)
		log.Println(dataHashed)

		// Send hashed data to Blockchain
		iotaHandler.StoreData(endpoint, seed, address, dataHashed, tag)

		t1 := time.Now()
		iotaHandler.RetriveData(endpoint, address, tag)
		t2 := time.Now()
		diff := t2.Sub(t1)
		log.Println(diff)
	}
}
