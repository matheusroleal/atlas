package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/iotaledger/iota.go/trinary"
	"github.com/matheusroleal/ozymandias/ozy/asset"
	iotaHandler "github.com/matheusroleal/ozymandias/ozy/blockchain/iota"
)

// const endpoint = "https://nodes.devnet.iota.org"
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
		// var tag = "XVLBZGBAICMRAJWWHTHCTCUAX"
		fmt.Println(tag)
		// Our data is very long here, it needs to be split over several transactions, 3 in this case
		sum := 0
		var bulk []string
		for i := 0; i < 100000; i++ { // Simulating 84 hours per week data
			data := asset.CreateHashData(fmt.Sprintf("%s%d", "data ", sum), fmt.Sprintf("%d", j))
			b, err := json.Marshal(data)
			if err != nil {
				panic(err)
			}
			dataParsed := string(b)
			// iotaHandler.StoreData(endpoint, seed, address, data, tag)
			bulk = append(bulk, dataParsed)
			// fmt.Println(i)
			sum += i
		}
		fmt.Println("Sending Messages in a Bulk")
		iotaHandler.BulkData(endpoint, seed, address, bulk, tag)
		t1 := time.Now()
		iotaHandler.RetriveData(endpoint, address, tag)
		t2 := time.Now()
		diff := t2.Sub(t1)
		fmt.Println(diff)
	}
}
