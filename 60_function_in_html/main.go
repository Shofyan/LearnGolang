package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type SuperHero struct {
	Name   string
	Alias  string
	Friend []string
}

func (s SuperHero) SayHello(From string, Message string) string {
	return fmt.Sprintf("%s said  \"%s\"", From, Message)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var super = SuperHero{
			Name:   "Bruce Wayne",
			Alias:  "Batman",
			Friend: []string{"Superman", "Wonder woman", "Aquaman"},
		}
		var tmpl = template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, super); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
