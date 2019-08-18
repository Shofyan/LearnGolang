// https://dasarpemrogramangolang.novalagung.com/48-web.html
package main

import (
	"fmt"
	"html/template"
)
import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hai dunia apakabar")
}

func displayMessage(w http.ResponseWriter, r *http.Request) {
	var data = map[string]string{
		"Name":    "Shofyan",
		"Message": "Apa kabar dunia, selamat pagi",
	}

	var t, err = template.ParseFiles("template.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	t.Execute(w, data)

}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "assalamualaikum ")
	})
	http.HandleFunc("/index", index)
	http.HandleFunc("/display", displayMessage)

	fmt.Println("starting webserver at http://localhost:8181")
	http.ListenAndServe(":8181", nil)
}
