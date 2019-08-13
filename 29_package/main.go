package main

import (
	"fmt"

	m1 "github.com/shofyan/LearnGolang/29_package/config"
	"github.com/shofyan/LearnGolang/29_package/model"
)

func main() {

	pc := m1.GetPostgresConnection()
	fmt.Println(pc)

	pp := model.Person{ID: 1, Name: "shofyan"}
	fmt.Println(pp)
}
