/*
 * @Author: Matheus Leal
 * @Date: 2022-07-01 22:54:49
 * @Last Modified by: Matheus Leal
 * @Last Modified time: 2022-07-03 14:08:24
 */
package iota

import (
	"bytes"
	"os"
	"sort"
	"strconv"
	"strings"

	. "github.com/iotaledger/iota.go/api"
	"github.com/iotaledger/iota.go/bundle"
	"github.com/iotaledger/iota.go/converter"
	"github.com/iotaledger/iota.go/trinary"
	log "github.com/sirupsen/logrus"
)

const mwm = 9
const depth = 3

/**
 * Store a new Asset into IOTA Blockchain.
 *
 * @param   endpoint 		string		IOTA endpoint string
 *					seed				string		IOTA seed string
 *					address			string		IOTA address string
 *					data				string		Asset data string
 *					tag					string		Asset identification string
 * @return  						Asset			A Asset struct
 */
func StoreData(endpoint string, seed string, address string, data string, tag string) error {
	// Compose a new API instance
	api, err := ComposeAPI(HTTPClientSettings{URI: endpoint})
	if err != nil {
		log.Error(err)
		return err
	}
	// Convert a ascii message for the transaction to trytes,if possible
	message, err := converter.ASCIIToTrytes(data)
	if err != nil {
		log.Error(err)
		return err
	}
	transfers := bundle.Transfers{
		{
			Address: address,
			Value:   0,
			Message: message,
			Tag:     trinary.Trytes(tag),
		},
	}
	// We need to pass an options object, since we want to use the defaults it stays empty
	prepTransferOpts := PrepareTransfersOptions{}
	trytes, err := api.PrepareTransfers(seed, transfers, prepTransferOpts)
	if err != nil {
		log.Error(err)
		return err
	}
	// Send the transaction to the tangle using given depth and minimum weight magnitude
	// _, err = api.SendTrytes(trytes, depth, mwm)
	bndl, err := api.SendTrytes(trytes, depth, mwm)
	if err != nil {
		log.Error(err)
		return err
	}
	var txhash = bundle.TailTransactionHash(bndl)
	log.Debug("[IOTA] broadcasted bundle with tail tx hash: " + txhash)
	log.Debug("[IOTA] https://explorer.iota.org/legacy-devnet/transaction/" + txhash)

	return nil
}

/**
 * Store a Bulk of Assets into IOTA Blockchain.
 *
 * @param   endpoint 		string		IOTA endpoint string
 *					seed				string		IOTA seed string
 *					address			string		IOTA address string
 *					bulk				string		A list of Asset data string
 *					tag					string		Asset identification string
 * @return  						Asset			A Asset struct
 */
func BulkData(endpoint string, seed string, address string, bulk []string, tag string) error {
	// Compose a new API instance
	api, err := ComposeAPI(HTTPClientSettings{URI: endpoint})
	if err != nil {
		log.Error(err)
		return err
	}
	var transfers bundle.Transfers
	limit, _ := strconv.ParseInt(os.Getenv("BULK_TRANSFER_LIMIT"), 10, 64)
	for i := 0; i < len(bulk); i += int(limit) {
		batch := bulk[i:min(i+int(limit), len(bulk))]
		transfers = prepareBulkArray(address, batch, tag)
		// We need to pass an options object, since we want to use the defaults it stays empty
		prepTransferOpts := PrepareTransfersOptions{}
		trytes, err := api.PrepareTransfers(seed, transfers, prepTransferOpts)
		if err != nil {
			log.Error(err)
			return err
		}
		// Send the transaction to the tangle using given depth and minimum weight magnitude
		bndl, err := api.SendTrytes(trytes, depth, mwm)
		// _, err = api.SendTrytes(trytes, depth, mwm)
		if err != nil {
			log.Error(err)
			return err
		}
		var txhash = bundle.TailTransactionHash(bndl)
		log.Debug("[IOTA] broadcasted bundle with tail tx hash: " + txhash)
		log.Debug("[IOTA] https://explorer.iota.org/legacy-devnet/transaction/" + txhash)
	}
	return nil
}

/**
 * Retrive a Asset from IOTA Blockchain.
 *
 * @param   endpoint 		string		IOTA endpoint string
 *					address			string		IOTA address string
 *					tag					string		Asset identification string
 * @return  						Asset			A Asset struct
 */
func RetriveData(endpoint string, address string, tag string) error {
	// Compose a new API instance
	api, err := ComposeAPI(HTTPClientSettings{URI: endpoint})
	if err != nil {
		log.Error(err)
		return err
	}
	// We need a query object containing the address we want to look for
	var modifiedTag = tag + strings.Repeat("9", 27-len(tag))
	var query = FindTransactionsQuery{Tags: trinary.Hashes{trinary.Trytes(modifiedTag)}}
	// Find Transaction Objects uses the connected node to find transactions based on our query
	transactions, err := api.FindTransactionObjects(query)
	if err != nil {
		log.Error(err)
		return err
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
	// We need to convert the message to ASCII, but before we do that we need to remove
	msg, err := converter.TrytesToASCII(removeSuffixNine(buffer.String()))
	if err != nil {
		log.Error(err)
		return err
	}
	// We print out the message
	log.Debug("[IOTA] DEBUG: Query Result " + msg)
	return nil
}

/**
 * Prepare a Bulk of Assets into a bundle.Transfers for IOTA Blockchain.
 *
 * @param   address		string						IOTA address string
 *					bulk			[]string					A list of Asset data string
 *					tag				string						Asset identification string
 * @return  					bundle.Transfers	A bundle.Transfers struct
 */
func prepareBulkArray(address string, bulk []string, tag string) bundle.Transfers {
	var transfers bundle.Transfers
	for _, data := range bulk {
		// Convert a ascii message for the transaction to trytes,if possible
		message, err := converter.ASCIIToTrytes(data)
		if err != nil {
			log.Error(err)
			return nil
		}
		// Create a transfer
		var transfer bundle.Transfer = bundle.Transfer{
			Address: address,
			Value:   0,
			Message: message,
			Tag:     trinary.Trytes(tag),
		}
		// Bulk multiple transfers to reduce time
		transfers = append(transfers, transfer)
	}
	return transfers
}

/**
 * Return the minumum value between 2 integers.
 *
 * @param   a		int		First integer
 *					b		int		Second integer
 * @return  		int		Minumum value between the 2 int
 */
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

/**
 * Remove the 9s suffixs from a string.
 *
 * @param   frag	string		A string containing 9s at the end
 * @return  			string		A string without 9s at the end
 */
func removeSuffixNine(frag string) string {
	fraglen := len(frag)
	var firstNonNineAt int
	for i := fraglen - 1; i > 0; i-- {
		if frag[i] != '9' {
			firstNonNineAt = i
			break
		}
	}
	return frag[:firstNonNineAt+1]
}
