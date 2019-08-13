package main

import (
	"flag"
	"fmt"
)

func main() {

	// config := flag.String("C", "1234", "your config")
	// flag.Parse()
	// myconfig := *config
	// fmt.Println("ini flag")
	// fmt.Println(myconfig)

	username := flag.String("U", "", "Your Username")
	password := flag.String("P", "", "Your Password")

	flag.Parse()
	result := login(*username, *password)

	if result {
		fmt.Println("login success")
	} else {
		fmt.Println("login failed")
	}
}

func login(username, password string) bool {
	if username == "shofyan" && password == "123456" {
		return true
	}
	return false

}
