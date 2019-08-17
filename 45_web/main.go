package main

import "fmt"
import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hai dunia apakabar")
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "assalamualaikum ")
	})
	http.HandleFunc("/index", index)

	fmt.Println("starting webserver at http://localhost:8080")
	http.ListenAndServe(":8181", nil)
}
