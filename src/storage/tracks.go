package storage

import (
	log "github.com/sirupsen/logrus"

	"github.com/matheusroleal/atlas/src/asset"
)

func IndexTracks(driver string, user string, password string, database string, address string) []asset.Asset {
	db := dbConn(driver, user, password, database, address)
	selDB, err := db.Query("SELECT * FROM Tracks ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := asset.Asset{}
	res := []asset.Asset{}
	for selDB.Next() {
		var id int
		var owner, data, reference string
		err = selDB.Scan(&id, &owner, &data, &reference)
		if err != nil {
			panic(err.Error())
		}
		emp.ID = string(rune(id))
		emp.Owner = owner
		emp.Data = data
		res = append(res, emp)
	}
	defer db.Close()
	return res
}

func ShowTrack(driver string, user string, password string, database string, address string, id string) []asset.Asset {
	db := dbConn(driver, user, password, database, address)
	selDB, err := db.Query("SELECT * FROM Tracks WHERE Reference=?", id)
	if err != nil {
		panic(err.Error())
	}
	res := []asset.Asset{}
	emp := asset.Asset{}
	for selDB.Next() {
		var id int
		var owner, data, reference string
		err = selDB.Scan(&id, &owner, &data, &reference)
		if err != nil {
			panic(err.Error())
		}
		emp.ID = string(rune(id))
		emp.Owner = owner
		emp.Data = data
		res = append(res, emp)
	}
	defer db.Close()
	return res
}

func InsertTrack(driver string, user string, password string, database string, address string, owner string, data string, reference string) {
	db := dbConn(driver, user, password, database, address)
	insForm, err := db.Prepare("INSERT INTO Tracks(owner, data, reference) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(owner, data, reference)
	log.Debug("[STORAGE] INSERT: Owner: " + owner + " | Data: " + data)
	defer db.Close()
}
