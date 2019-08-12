package main

import (
	"fmt"
	"time"
)

func main() {

	go helloGo()

	time.Sleep(1 * time.Second)
	fmt.Println("hai semua")
}

func helloGo() {
	fmt.Println("hello semua")
}
