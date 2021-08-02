package iota

import (
  . "github.com/iotaledger/iota.go/api"
  "github.com/iotaledger/iota.go/trinary"
  "github.com/iotaledger/iota.go/converter"
  "github.com/iotaledger/iota.go/bundle"
  "strings"
  // "log"
  "bytes"
  "sort"
)

const mwm = 9
const depth = 3

func StoreData(endpoint string, seed string, address string, data string, tag string) {
  // Compose a new API instance
  api, err := ComposeAPI(HTTPClientSettings{URI: endpoint})
  if err != nil {
    panic(err)
  }
  // Convert a ascii message for the transaction to trytes,if possible
  message, err := converter.ASCIIToTrytes(data)
  if err != nil {
    panic(err)
  }
  transfers := bundle.Transfers{
    {
      Address: address,
      Value: 0,
      Message: message,
      Tag: trinary.Trytes(tag),
    },
  }
  // We need to pass an options object, since we want to use the defaults it stays empty
  prepTransferOpts := PrepareTransfersOptions{}
  trytes, err := api.PrepareTransfers(seed, transfers, prepTransferOpts)
  if err != nil {
    panic(err)
  }
  // Send the transaction to the tangle using given depth and minimum weight magnitude
  _, err = api.SendTrytes(trytes, depth, mwm)
  // bndl, err := api.SendTrytes(trytes, depth, mwm)
  if err != nil {
    panic(err)
  }
  // var txhash = bundle.TailTransactionHash(bndl)
  // log.Println("[IOTA] DEBUG: broadcasted bundle with tail tx hash: ", txhash)
  // log.Printf("[IOTA] DEBUG: https://explorer.iota.org/legacy-devnet/transaction/%s\n\n", txhash)
}

func RetriveData(endpoint string, address string, tag string){
  // Compose a new API instance
  api, err := ComposeAPI(HTTPClientSettings{URI: endpoint})
  if err != nil {
    panic(err) 
  }
  // We need a query object containing the address we want to look for
  var modifiedTag = tag + strings.Repeat("9", 27 - len(tag))
  var query = FindTransactionsQuery{Tags: trinary.Hashes{trinary.Trytes(modifiedTag)}}
  // Find Transaction Objects uses the connected node to find transactions based on our query
  transactions, err := api.FindTransactionObjects(query)
  if err != nil {
    panic(err) 
  }
  // We need to sort all transactions by index first so we can concatenate them
  sort.Slice(transactions[:], func(i, j int) bool {
    return transactions[i].CurrentIndex < transactions[j].CurrentIndex
  })
  // We define a buffer to concatenate all sorted transactions
  var buffer bytes.Buffer
  // We add the sorted Transaction Signature Message Fragment to the buffer
  for _, tx := range transactions {
    buffer.WriteString(tx.SignatureMessageFragment)
  }
  // // We need to convert the message to ASCII, but before we do that we need to remove 
  // msg, err := converter.TrytesToASCII(removeSuffixNine(buffer.String()))
  // if err != nil {
  //   panic(err) 
  // }
  // // We print out the message
  // log.Println("[IOTA] DEBUG: Query Result ", msg)
}

func removeSuffixNine(frag string) string {
  fraglen := len(frag)
  var firstNonNineAt int
  for i := fraglen - 1; i > 0; i-- {
    if frag[i] != '9' {
      firstNonNineAt = i
      break;
    }
  }
  return frag[:firstNonNineAt+1]
}