/*
 * @Author: Matheus Leal
 * @Date: 2022-07-01 22:54:54
 * @Last Modified by: Matheus Leal
 * @Last Modified time: 2022-07-03 12:38:55
 */
package asset

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateAsset(t *testing.T) {
	tests := []struct {
		context  string
		mobileID string
		ref      string
	}{
		{"context1", "1", "reference1"},
		{"context2", "1", "reference2"},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("TestCreateAsset=%s", tc.context), func(t *testing.T) {
			got := CreateAsset(tc.context, tc.mobileID, tc.ref)
			if got != nil {
				t.Logf("Success !")
			} else {
				t.Fatalf("got %v", got)
			}
		})
	}
}

func TestCompressAsset(t *testing.T) {
	tests := []struct {
		asset []Asset
	}{
		{},
		{},
	}
	for _, tc := range tests {
		tc.asset = append(tc.asset, *CreateAsset("context1", "1", "reference1"))
		tc.asset = append(tc.asset, *CreateAsset("context2", "1", "reference2"))
		t.Run(fmt.Sprintf("TestCompressAsset=%s", tc.asset), func(t *testing.T) {
			got := CompressAsset(tc.asset)
			if reflect.TypeOf(got) == reflect.TypeOf("string") {
				t.Logf("Success !")
			} else {
				t.Fatalf("got %v", got)
			}
		})
	}
}

func TestHashAsset(t *testing.T) {
	tests := []struct {
		data string
	}{
		{"context1"},
		{"context2"},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("TestHashAsset=%s", tc.data), func(t *testing.T) {
			got := HashAsset(tc.data)
			if reflect.TypeOf(got) == reflect.TypeOf("string") {
				t.Logf("Success !")
			} else {
				t.Fatalf("got %v", got)
			}
		})
	}
}
