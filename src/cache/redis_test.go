/*
 * @Author: Matheus Leal
 * @Date: 2022-07-01 22:54:39
 * @Last Modified by: Matheus Leal
 * @Last Modified time: 2022-07-02 12:22:20
 */
package cache

import (
	"fmt"
	"testing"
)

func TestCacheConn(t *testing.T) {
	tests := []struct {
		address  string
		password string
	}{
		{"0.0.0.0:6379", ""},
		{"0.0.0.0:6379", ""},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("TestcacheConn=%s", tc.address), func(t *testing.T) {
			got := cacheConn(tc.address, tc.password)
			if got != nil {
				t.Logf("Success !")
			} else {
				t.Fatalf("got %v", got)
			}
		})
	}
}

func TestGetData(t *testing.T) {
	tests := []struct {
		address  string
		password string
		key      string
		want     string
	}{
		{"0.0.0.0:6379", "", "test", ""},
		{"0.0.0.0:6379", "", "test2", "123"},
		{"0.0.0.0:6379", "", "test3", "segment"},
		{"0.0.0.0:6379", "", "test4", ""},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("TestGetData=%s", tc.key), func(t *testing.T) {
			_ = SetData(tc.address, tc.password, tc.key, tc.want)
			got := GetData(tc.address, tc.password, tc.key)
			if got == tc.want {
				t.Logf("Success !")
			} else {
				t.Fatalf("got %v", got)
			}
		})
	}
}

func TestSetData(t *testing.T) {
	tests := []struct {
		address  string
		password string
		key      string
		want     error
	}{
		{"0.0.0.0:6379", "", "test5", nil},
		{"0.0.0.0:6379", "", "test6", nil},
		{"0.0.0.0:6379", "", "test7", nil},
		{"0.0.0.0:6379", "", "test8", nil},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("TestSetData=%s", tc.key), func(t *testing.T) {
			got := SetData(tc.address, tc.password, tc.key, tc.key)
			if got == tc.want {
				t.Logf("Success !")
			} else {
				t.Fatalf("got %v", got)
			}
		})
	}
}

func TestAppendData(t *testing.T) {
	tests := []struct {
		address  string
		password string
		key      string
		want     error
	}{
		{"0.0.0.0:6379", "", "test5", nil},
		{"0.0.0.0:6379", "", "test6", nil},
		{"0.0.0.0:6379", "", "test7", nil},
		{"0.0.0.0:6379", "", "test8", nil},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("TestAppendData=%s", tc.key), func(t *testing.T) {
			got := AppendData(tc.address, tc.password, tc.key, tc.key+",")
			if got == tc.want {
				t.Logf("Success !")
			} else {
				t.Fatalf("got %v", got)
			}
		})
	}
}
