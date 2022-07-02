/*
 * @Author: Matheus Leal
 * @Date: 2022-07-01 22:54:44
 * @Last Modified by: Matheus Leal
 * @Last Modified time: 2022-07-02 11:28:17
 */
package iota

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNodeInfo(t *testing.T) {
	const endpoint = "https://nodes.devnet.iota.org:443"
	tests := []struct {
		test int
	}{
		{1},
		{2},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("TestNodeInfo=%d", tc.test), func(t *testing.T) {
			got := NodeInfo(endpoint)
			if got == nil {
				t.Logf("Success !")
			} else {
				t.Fatalf("got %v", got)
			}
		})
	}
}

func TestRandStringRunes(t *testing.T) {
	const endpoint = "https://nodes.devnet.iota.org:443"
	tests := []struct {
		n    int
		want string
	}{
		{1, "string"},
		{2, "string"},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("TestStoreData=%d", tc.n), func(t *testing.T) {
			got := RandStringRunes(tc.n)
			if reflect.TypeOf(got) == reflect.TypeOf(tc.want) {
				t.Logf("Success !")
			} else {
				t.Fatalf("got %v", got)
			}
		})
	}
}
