/*
 * @Author: Matheus Leal
 * @Date: 2022-07-01 22:54:18
 * @Last Modified by: Matheus Leal
 * @Last Modified time: 2022-07-03 11:39:22
 */
package storage

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDBConn(t *testing.T) {
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
		t.Run(fmt.Sprintf("TestDBConn=%s", tc.address), func(t *testing.T) {
			got := dbConn(tc.driver, tc.user, tc.password, tc.database, tc.address)
			if got != nil {
				t.Logf("Success !")
			} else {
				t.Fatalf("got %v", got)
			}
		})
	}
}
