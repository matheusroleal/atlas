
import (
	"github.com/matheusroleal/ozymandias/ozy/asset"
)

func IndexSegments(driver string, user string, password string, database string) {
	db := dbConn(driver, user, password, database)
	selDB, err := db.Query("SELECT * FROM Segment ORDER BY id DESC")
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
		emp.Id = id
		emp.Owner = owner
		emp.Data = data
		res = append(res, emp)
	}
	defer db.Close()
}

func ShowSegment(driver string, user string, password string, database string, id string) {
	db := dbConn(driver, user, password, database)
	selDB, err := db.Query("SELECT * FROM Segment WHERE id=?", id)
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
		emp.Id = id
		emp.Owner = owner
		emp.Data = data
	}
	defer db.Close()
}

func InsertSegment(driver string, user string, password string, database string, id string, owner string, data string) {
	db := dbConn(driver, user, password, database)
	insForm, err := db.Prepare("INSERT INTO Segment(id, owner, data) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(id, owner, data)
	log.Println("INSERT: Owner: " + owner + " | Data: " + data)
	defer db.Close()
}
