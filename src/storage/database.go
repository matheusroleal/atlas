/*
 * @Author: Matheus Leal
 * @Date: 2022-07-01 22:54:18
 * @Last Modified by: Matheus Leal
 * @Last Modified time: 2022-07-03 14:17:01
 */
package storage

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/**
 * Start a new connection with a database server.
 *
 * @param   driver			string			A database driver
 *					dbUser			string	 		A database user
 *					dbPass			string	 		A database password
 *					dbName			string	 		A database name
 *					dbAddress		string	 		A database address
 * @return  db					sql.DB			A SQL client
 */
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
