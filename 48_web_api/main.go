//https://dasarpemrogramangolang.novalagung.com/51-web-json-api.html

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type student struct {
	ID    string
	Name  string
	Grade int
}

var data = []student{
	{"a01", "shofyan", 90},
	{"a02", "sky", 78},
	{"a03", "pencari", 80},
	{"a04", "bear", 60},
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var result, err = json.Marshal(data)
		if err != nil {
			fmt.Println(err.Error())
		}

		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var id = r.FormValue("id")
		var result []byte
		var err error

		for _, each := range data {
			if each.ID == id {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

				w.Write(result)
				return

			}
		}

		http.Error(w, "user not fount", http.StatusBadRequest)
		return
	}
}

func main() {
	http.HandleFunc("/user", user)
	http.HandleFunc("/users", users)

	fmt.Println("starting web server at http://localhost:8181")
	http.ListenAndServe(":8181", nil)
}
