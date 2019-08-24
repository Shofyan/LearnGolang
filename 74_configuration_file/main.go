package main

import (
	"fmt"
	conf "github.com/Shofyan/LearnGolang/74_configuration_file/config"
	"log"
	"net/http"
)

type CustomMux struct {
	http.ServeMux
}

func (c CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if conf.Configuration().Log.Verbose {
		log.Println("Incoming request from ", r.Host, "accessing", r.URL.String())
	}
	c.ServeMux.ServeHTTP(w, r)
}

func main() {
	router := new(CustomMux)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello dunia"))
	})
	router.HandleFunc("/howareyou", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("how are you?"))
	})

	server := new(http.Server)
	server.Handler = router
	server.ReadTimeout = conf.Configuration().Server.ReadTimeout
	server.WriteTimeout = conf.Configuration().Server.WriteTimeout
	server.Addr = fmt.Sprintf(":%d", conf.Configuration().Server.Port)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
		return
	}

}
