package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Info struct {
	Affiliation string
	Address     string
}

type Person struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    Info
}

func (t Info) GetAffiliationDetailInfo() string {
	return "Have 31 division"
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		person := Person{
			Name:    "Si Buta Goa Hantu",
			Gender:  "Pria",
			Hobbies: []string{"Bertarung", "semedi", "berkelana"},
			Info:    Info{"Goa tak berpenghuni", "jakarta barat"},
		}

		var tmpl = template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Print("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)

}
