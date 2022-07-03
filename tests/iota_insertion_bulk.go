/*
 * @Author: Matheus Leal
 * @Date: 2022-07-01 22:57:46
 * @Last Modified by: Matheus Leal
 * @Last Modified time: 2022-07-03 12:08:06
 */
package test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/matheusroleal/atlas/src/asset"
	iotaHandler "github.com/matheusroleal/atlas/src/blockchain/iota"
	log "github.com/sirupsen/logrus"
)

func bulkInsertionTest() {
	// Create blocks of data with tags
	for j := 0; j < 31; j++ {
		var tag = iotaHandler.RandStringRunes(25)
		log.Debug("[Atlas][Test] " + tag)
		// Our data is very long here, it needs to be split over several transactions, 3 in this case
		sum := 0
		var segments []string

		for i := 0; i < 500; i++ {
			data := asset.CreateAsset(fmt.Sprintf("%s%d", "checkpoint ", sum), fmt.Sprintf("%d", j), "Track1")

			// Convert asset to a hash
			b, err := json.Marshal(data)
			if err != nil {
				return
			}
			dataParsed := string(b)
			dataHashed := asset.HashAsset(dataParsed)

			// Bulk hashed data to be sent to Blockchain
			segments = append(segments, dataHashed)
			sum += i
		}

		t1 := time.Now()
		// Send hashed data to Blockchain
		iotaHandler.BulkData(endpoint, seed, address, segments, tag)
		t2 := time.Now()
		diff := t2.Sub(t1)
		log.Info(diff)
	}
}
