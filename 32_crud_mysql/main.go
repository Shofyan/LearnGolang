package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

//Employee adalah contoh domain yang akan disave
type Employee struct {
	ID   int
	Name string
	City string
}

func dbcond() (db *sql.DB) {
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

// tmpl variabel global untuk diakses tiap method
var tmpl = template.Must(template.ParseGlob("form/*"))

// Index untuk akses web pertama kali
func Index(w http.ResponseWriter, r *http.Request) {
	db := dbcond()

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
		log.Println(emp)
	}

	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

// Show ialah untuk menampilkan halaman detail
func Show(w http.ResponseWriter, r *http.Request) {
	db := dbcond()
	nID := r.URL.Query().Get("id")
	selDb, err := db.Query("SELECT * FROM Employee where id=?", nID)
	if err != nil {
		panic(err.Error())
	}

	emp := Employee{}
	for selDb.Next() {
		var id int
		var name, city string
		err = selDb.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.ID = id
		emp.Name = name
		emp.City = city

	}

	tmpl.ExecuteTemplate(w, "Show", emp)
	defer db.Close()

}

// New untuk menampilkan halaman create awal
func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

// Edit menampilkan halaman input yang sudah ada dtanya
func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbcond()
	nID := r.URL.Query().Get("id")
	selDb, err := db.Query("SELECT * FROM Employee where id=?", nID)
	if err != nil {
		panic(err.Error())
	}

	emp := Employee{}
	for selDb.Next() {
		var id int
		var name, city string
		err = selDb.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.ID = id
		emp.Name = name
		emp.City = city

	}
	defer db.Close()

	tmpl.ExecuteTemplate(w, "Edit", emp)

}

// Insert data employee . post method
func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbcond()

	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		insForm, err := db.Prepare("INSERT INTO Employee(name,city) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}

		insForm.Exec(name, city)
		log.Println("INSERT: NAME: " + name + " | CITY: " + city)
	}
	defer db.Close()

	http.Redirect(w, r, "/", 301)

}

// Update untuk update data ketika submit
func Update(w http.ResponseWriter, r *http.Request) {
	db := dbcond()

	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		uID := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE Employee SET name=?, city=? Where id=?")
		if err != nil {
			panic(err.Error())
		}

		insForm.Exec(name, city, uID)
		log.Println("UPDATE: NAME: " + name + " | CITY: " + city)
	}
	defer db.Close()

	http.Redirect(w, r, "/", 301)

}

// Delete untuk Delete data
func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbcond()

	ID := r.URL.Query().Get("id")
	insForm, err := db.Prepare("DELETE from Employee Where id=?")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(ID)
	log.Println("Delete")
	defer db.Close()

	http.Redirect(w, r, "/", 301)

}

func main() {
	log.Println("Server started on : http://Localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}
