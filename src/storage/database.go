package storage

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn(dbDriver string, dbUser string, dbPass string, dbName string, dbAddress string) (db *sql.DB) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbAddress+"/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	// Important settings
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
