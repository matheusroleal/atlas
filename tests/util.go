/*
 * @Author: Matheus Leal
 * @Date: 2022-07-01 22:57:56
 * @Last Modified by: Matheus Leal
 * @Last Modified time: 2022-07-01 22:59:52
 */
package test

import (
	"math/rand"

	"github.com/iotaledger/iota.go/trinary"
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
