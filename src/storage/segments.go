/*
 * @Author: Matheus Leal
 * @Date: 2022-07-01 22:54:22
 * @Last Modified by: Matheus Leal
 * @Last Modified time: 2022-07-03 14:19:48
 */
package storage

import (
	log "github.com/sirupsen/logrus"

	"github.com/matheusroleal/atlas/src/asset"
)

/**
 * Get the segment index from database server.
 *
 * @param   driver			string			A database driver
 *					dbUser			string	 		A database user
 *					dbPass			string	 		A database password
 *					dbName			string	 		A database name
 *					dbAddress		string	 		A database address
 * @return  db					sql.DB			A SQL client
 */
func IndexSegments(driver string, user string, password string, database string, address string) ([]asset.Asset, error) {
	db := dbConn(driver, user, password, database, address)
	selDB, err := db.Query("SELECT * FROM Segments ORDER BY id DESC")
	if err != nil {
		log.Error(err)
		return nil, err
	}
	emp := asset.Asset{}
	res := []asset.Asset{}
	for selDB.Next() {
		var id int
		var owner, data, reference string
		err = selDB.Scan(&id, &owner, &data, &reference)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		emp.ID = string(rune(id))
		emp.Owner = owner
		emp.Data = data
		res = append(res, emp)
	}
	defer db.Close()
	return res, nil
}

/**
 * Get a segment data from database server.
 *
 * @param   driver			string			A database driver
 *					dbUser			string	 		A database user
 *					dbPass			string	 		A database password
 *					dbName			string	 		A database name
 *					dbAddress		string	 		A database address
 *					id					string	 		A segment identification
 * @return  db					sql.DB			A SQL client
 */
func GetSegment(driver string, user string, password string, database string, address string, id string) ([]asset.Asset, error) {
	db := dbConn(driver, user, password, database, address)
	selDB, err := db.Query("SELECT * FROM Segments WHERE Reference=?", id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	res := []asset.Asset{}
	emp := asset.Asset{}
	for selDB.Next() {
		var id int
		var owner, data, reference string
		err = selDB.Scan(&id, &owner, &data, &reference)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		emp.ID = string(rune(id))
		emp.Owner = owner
		emp.Data = data
		res = append(res, emp)
	}
	defer db.Close()
	return res, nil
}

/**
 * Insert a segment data from database server.
 *
 * @param   driver			string			A database driver
 *					dbUser			string	 		A database user
 *					dbPass			string	 		A database password
 *					dbName			string	 		A database name
 *					dbAddress		string	 		A database address
 *					owner				string	 		A segment owner
 *					data				string	 		A segment data
 *					reference		string	 		A segment reference
 * @return  db					sql.DB			A SQL client
 */
func InsertSegment(driver string, user string, password string, database string, address string, owner string, data string, reference string) error {
	db := dbConn(driver, user, password, database, address)
	insForm, err := db.Prepare("INSERT INTO Segments(owner, data, reference) VALUES(?,?,?)")
	if err != nil {
		log.Error(err)
		return err
	}
	insForm.Exec(owner, data, reference)
	log.Debug("[STORAGE] INSERT: Owner: " + owner + " | Data: " + data)
	defer db.Close()
	return nil
}
