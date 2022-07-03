/*
 * @Author: Matheus Leal
 * @Date: 2022-07-01 22:57:56
 * @Last Modified by: Matheus Leal
 * @Last Modified time: 2022-07-02 10:24:37
 */
package test

import (
	"github.com/iotaledger/iota.go/trinary"
)

const endpoint = "https://nodes.devnet.iota.org:443"

// We need a dummy seed even though we don't sign, because the API requires a seed to send
const seed = trinary.Trytes("JBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQ")
const address = trinary.Trytes("XBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQMZOCAHYPD")
