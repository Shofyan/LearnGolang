package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", actionIndex)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func actionIndex(w http.ResponseWriter, r *http.Request) {
	data := []struct {
		Name string
		Age  int
	}{
		{"shofyan", 23},
		{"wahyu", 40},
		{"ilham", 23},
		{"sary", 30},
	}

	jsonByte, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonByte)

}
