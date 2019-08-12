package main

import (
	"fmt"
)

func main() {

	done := make(chan bool)

	go helloGo(done)
	<-done

	fmt.Println("hai semua")
}

func helloGo(done chan bool) {
	fmt.Println("hello semua")
	done <- true
}
