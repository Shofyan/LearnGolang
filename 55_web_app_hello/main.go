package main

import "fmt"
import "net/http"

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Welcome"
	_, _ = w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "hello world"
	_, _ = w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	var address = "localhost:9000"
	fmt.Println("server started as %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}
