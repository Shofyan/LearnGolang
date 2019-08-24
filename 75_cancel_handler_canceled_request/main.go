package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	done := make(chan bool)
	go func() {
		body, _ := ioutil.ReadAll(r.Body)
		log.Println(body)

		time.Sleep(time.Second * 10)
		done <- true
	}()

	select {
	case <-r.Context().Done():
		if err := r.Context().Err(); err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "canceled") {
				log.Println("request canceled")
			} else {
				log.Println("unknown error ocured", err.Error())
			}
		}
	case <-done:
		log.Println("done")
	}
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":8080", nil)
}
