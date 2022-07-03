/*
 * @Author: Matheus Leal
 * @Date: 2022-07-01 22:54:22
 * @Last Modified by: Matheus Leal
 * @Last Modified time: 2022-07-03 12:17:57
 */
package storage

import (
	"fmt"
	"testing"
)

func TestIndexTracks(t *testing.T) {
	tests := []struct {
		driver   string
		user     string
		password string
		database string
		address  string
	}{
		{"mysql", "root", "password", "Atlas", "tcp(0.0.0.0:6603)"},
		{"mysql", "root", "password", "Atlas", "tcp(0.0.0.0:6603)"},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("TestIndexTracks=%s", tc.address), func(t *testing.T) {
			_, got := IndexTracks(tc.driver, tc.user, tc.password, tc.database, tc.address)
			if got == nil {
				t.Logf("Success !")
			} else {
				t.Fatalf("got %v", got)
			}
		})
	}
}

func TestGetTrack(t *testing.T) {
	tests := []struct {
		driver   string
		user     string
		password string
		database string
		address  string
		id       string
	}{
		{"mysql", "root", "password", "Atlas", "tcp(0.0.0.0:6603)", "1"},
		{"mysql", "root", "password", "Atlas", "tcp(0.0.0.0:6603)", "1"},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("TestIndexTracks=%s", tc.address), func(t *testing.T) {
			_, got := GetTrack(tc.driver, tc.user, tc.password, tc.database, tc.address, tc.id)
			if got == nil {
				t.Logf("Success !")
			} else {
				t.Fatalf("got %v", got)
			}
		})
	}
}

func TestInsertTrack(t *testing.T) {
	tests := []struct {
		driver    string
		user      string
		password  string
		database  string
		address   string
		owner     string
		data      string
		reference string
	}{
		{"mysql", "root", "password", "Atlas", "tcp(0.0.0.0:6603)", "1", "track1", "Track1"},
		{"mysql", "root", "password", "Atlas", "tcp(0.0.0.0:6603)", "1", "track1", "Track1"},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("TestIndexTracks=%s", tc.address), func(t *testing.T) {
			got := InsertTrack(tc.driver, tc.user, tc.password, tc.database, tc.address, tc.owner, tc.data, tc.reference)
			if got == nil {
				t.Logf("Success !")
			} else {
				t.Fatalf("got %v", got)
			}
		})
	}
}
