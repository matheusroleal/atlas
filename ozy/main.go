package main

import (
	iotaHandler "github.com/matheusroleal/ozymandias/ozy/iota"
  asset "github.com/matheusroleal/ozymandias/ozy/asset"
	"github.com/iotaledger/iota.go/trinary"
  "encoding/json"
  "time"
  "fmt"
  "math/rand"
)

const endpoint = "https://nodes.devnet.iota.org"
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
  for j := 0; j < 43800; j++ {
    var tag = RandStringRunes(25)
    // Our data is very long here, it needs to be split over several transactions, 3 in this case
    sum := 0
    for i := 0; i < 2628002; i++ {
      checkpoint := asset.CreateCheckpoint(fmt.Sprintf("%s%d","checkpoint ", sum), fmt.Sprintf("%d",j))
      b, err := json.Marshal(checkpoint)
      if err != nil {
          panic(err)
      }
      data := string(b)
      iotaHandler.StoreData(endpoint, seed, address, data, tag)
      sum += i
    }
    t1 := time.Now()
    iotaHandler.RetriveData(endpoint, address, tag)
    t2 := time.Now()
    diff := t2.Sub(t1)
    fmt.Println(diff)
  }
}