// referece article https://medium.com/@kiddy.xyz/tutorial-golang-rest-api-mysql-part-1-45cd9f4e75a6

package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//Employee adalah contoh domain yang akan disave
type Employee struct {
	ID   int    `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
	City string `form:"city" json:"city"`
}

// ResponseEmployee cuma untuk standarasasi balikan json
type ResponseEmployee struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Employee
}

func dbconn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "goblog"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// ReturnAllEmpmployees ngembaliin semua employee
func ReturnAllEmpmployees(w http.ResponseWriter, r *http.Request) {
	db := dbconn()
	defer db.Close()

	selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	emp := Employee{}
	res := []Employee{}

	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}

		emp.ID = id
		emp.Name = name
		emp.City = city
		res = append(res, emp)
	}

	rspnd := ResponseEmployee{
		Status:  1,
		Message: "Success",
		Data:    res,
	}

	w.Header().Set("Contet-Type", "application/json")
	json.NewEncoder(w).Encode(rspnd)

}

func main() {
	log.Println("Server started on : http://Localhost:8080")
	http.HandleFunc("/", ReturnAllEmpmployees)

	http.ListenAndServe(":8080", nil)
}
