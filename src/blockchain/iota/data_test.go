/*
 * @Author: Matheus Leal
 * @Date: 2022-07-02 10:14:19
 * @Last Modified by: Matheus Leal
 * @Last Modified time: 2022-07-02 11:24:23
 */
package iota

import (
	"fmt"
	"testing"

	"github.com/iotaledger/iota.go/trinary"

	"github.com/matheusroleal/atlas/src/asset"
)

func TestStoreData(t *testing.T) {
	const endpoint = "https://nodes.devnet.iota.org:443"
	const seed = trinary.Trytes("JBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQ")
	const address = trinary.Trytes("XBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQMZOCAHYPD")
	tests := []struct {
		data string
		want error
	}{
		{"12738efdfb4fd30a3bc9b9ed725cc42c512100510f3745bb1a2d2c8a5b752c54", nil},
		{"b4db3230dc0fc2ec7cbf064574943aeb1baaf0227c7505a27878a2d22f4b2333", nil},
		{"617ab5d33ab15567cafa8e8e32a403b5ccee53eb839e50e44a9e95410891fa1d", nil},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("TestStoreData=%s", tc.data), func(t *testing.T) {
			var tag = RandStringRunes(25)
			dataHashed := asset.HashAsset(tc.data)
			got := StoreData(endpoint, seed, address, dataHashed, tag)
			if got == nil {
				t.Logf("Success !")
			} else {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}

func TestBulkData(t *testing.T) {
	const endpoint = "https://nodes.devnet.iota.org:443"
	const seed = trinary.Trytes("JBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQ")
	const address = trinary.Trytes("XBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQMZOCAHYPD")
	tests := []struct {
		test int
	}{
		{1},
		{2},
	}
	for _, tc := range tests {
		var data = []string{"10", "50", "30", "40", "50"}
		t.Run(fmt.Sprintf("TestBulkData=%d", tc.test), func(t *testing.T) {
			var tag = RandStringRunes(25)
			got := BulkData(endpoint, seed, address, data, tag)
			if got == nil {
				t.Logf("Success !")
			} else {
				t.Fatalf("got %v", got)
			}
		})
	}
}

func TestRetriveData(t *testing.T) {
	const endpoint = "https://nodes.devnet.iota.org:443"
	const seed = trinary.Trytes("JBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQ")
	const address = trinary.Trytes("XBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQMZOCAHYPD")
	tests := []struct {
		data string
		want error
	}{
		{"12738efdfb4fd30a3bc9b9ed725cc42c512100510f3745bb1a2d2c8a5b752c54", nil},
		{"b4db3230dc0fc2ec7cbf064574943aeb1baaf0227c7505a27878a2d22f4b2333", nil},
		{"617ab5d33ab15567cafa8e8e32a403b5ccee53eb839e50e44a9e95410891fa1d", nil},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("TestRetriveData=%s", tc.data), func(t *testing.T) {
			var tag = RandStringRunes(25)
			dataHashed := asset.HashAsset(tc.data)
			StoreData(endpoint, seed, address, dataHashed, tag)
			got := RetriveData(endpoint, address, tag)
			if got == tc.want {
				t.Logf("Success !")
			} else {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}

func TestPrepareBulkArray(t *testing.T) {
	const address = trinary.Trytes("XBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQMZOCAHYPD")
	tests := []struct {
		test int
	}{
		{1},
		{2},
	}
	for _, tc := range tests {
		var data = []string{"10", "50", "30", "40", "50"}
		t.Run(fmt.Sprintf("TestPrepareBulkArray=%d", tc.test), func(t *testing.T) {
			var tag = RandStringRunes(25)
			got := prepareBulkArray(address, data, tag)
			if got != nil {
				t.Logf("Success !")
			} else {
				t.Fatalf("got %v", got)
			}
		})
	}
}

func TestRemoveSuffixNine(t *testing.T) {
	tests := []struct {
		frag string
		want string
	}{
		{"12341234999", "12341234"},
		{"1234", "1234"},
		{"129349", "12934"},
		{"12394", "12394"},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("TestPrepareBulkArray=%s", tc.frag), func(t *testing.T) {
			got := removeSuffixNine(tc.frag)
			if got == tc.want {
				t.Logf("Success !")
			} else {
				t.Fatalf("got %v", got)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{1, 2, 1},
		{2, 1, 1},
		{5, 2, 2},
		{0, 2, 0},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("TestMin=%d and %d", tc.a, tc.b), func(t *testing.T) {
			got := min(tc.a, tc.b)
			if got == tc.want {
				t.Logf("Success !")
			} else {
				t.Fatalf("got %v", got)
			}
		})
	}
}
