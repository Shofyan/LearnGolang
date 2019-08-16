// referece article https://medium.com/@kiddy.xyz/tutorial-golang-rest-api-mysql-part-1-45cd9f4e75a6

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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

// ReturnAllEmployees ngembaliin semua employee
func ReturnAllEmployees(w http.ResponseWriter, r *http.Request) {
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

	response := ResponseEmployee{
		Status:  1,
		Message: "Success",
		Data:    res,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func insertEmployeeMultipart(w http.ResponseWriter, r *http.Request) {

	var response ResponseEmployee

	db := dbconn()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	name := r.FormValue("Name")
	city := r.FormValue("CIty")

	_, err = db.Exec("INSERT INTO EMPLOYEE (Name,City) values(?,?)", name, city)

	if err != nil {
		log.Println(err)
	}

	response.Status = 1
	response.Message = "Success Add"
	log.Println("insert data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func updateEmployeeMultipart(w http.ResponseWriter, r *http.Request) {
	db := dbconn()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	id := r.FormValue("ID")
	name := r.FormValue("Name")
	city := r.FormValue("City")

	_, err = db.Exec("update Employee set name=?,city=? where id = ?", name, city, id)
	if err != nil {
		log.Print(err)
	}

	response := ResponseEmployee{
		Status:  1,
		Message: "Success Update Data",
	}
	log.Println("Update data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func deleteEmployeesMultipart(w http.ResponseWriter, r *http.Request) {
	db := dbconn()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	id := r.FormValue("ID")
	_, err = db.Exec("DELETE FROM EMPLOYEE where id=?", id)
	if err != nil {
		log.Print(err)
	}
	response := ResponseEmployee{
		Status:  1,
		Message: "Success Delete Data",
	}
	log.Println("Delete data to database")

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/ReturnAllEmployees", ReturnAllEmployees).Methods("GET")
	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

	log.Println("Server started on : http://Localhost:8080")
	http.HandleFunc("/", ReturnAllEmployees)
	router.HandleFunc("/employees", insertEmployeeMultipart).Methods("POST")
	router.HandleFunc("/employees", updateEmployeeMultipart).Methods("PUT")
	router.HandleFunc("/users", deleteEmployeesMultipart).Methods("DELETE")

	http.ListenAndServe(":8080", nil)
}
