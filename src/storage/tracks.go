package storage

import (
	"log"

	"github.com/matheusroleal/atlas/src/asset"
)

func IndexTracks(driver string, user string, password string, database string, address string) {
	db := dbConn(driver, user, password, database, address)
	selDB, err := db.Query("SELECT * FROM Tracks ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := asset.Asset{}
	res := []asset.Asset{}
	for selDB.Next() {
		var id int
		var owner, data string
		err = selDB.Scan(&id, &owner, &data)
		if err != nil {
			panic(err.Error())
		}
		emp.ID = string(rune(id))
		emp.Owner = owner
		emp.Data = data
		res = append(res, emp)
	}
	defer db.Close()
}

func ShowTrack(driver string, user string, password string, database string, address string, id string) {
	db := dbConn(driver, user, password, database, address)
	selDB, err := db.Query("SELECT * FROM Tracks WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}
	emp := asset.Asset{}
	for selDB.Next() {
		var id int
		var owner, data string
		err = selDB.Scan(&id, &owner, &data)
		if err != nil {
			panic(err.Error())
		}
		emp.ID = string(rune(id))
		emp.Owner = owner
		emp.Data = data
	}
	defer db.Close()
}

func InsertTrack(driver string, user string, password string, database string, address string, owner string, data string) {
	db := dbConn(driver, user, password, database, address)
	insForm, err := db.Prepare("INSERT INTO Tracks(owner, data) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(owner, data)
	log.Println("INSERT: Owner: " + owner + " | Data: " + data)
	defer db.Close()
}