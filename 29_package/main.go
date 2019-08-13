package main

import (
	"fmt"

	"./config"
	"./model"
)

func main() {

	pc := config.GetPostgresConnection()
	fmt.Println(pc)

	pp := model.Person{ID: 1, Name: "shofyan"}
	fmt.Println(pp)
}
