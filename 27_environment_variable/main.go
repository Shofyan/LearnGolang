package main

import (
	"fmt"
	"os"
)

func main() {

	os.Setenv("POSTGRES", "posgres.host")
	os.Setenv("MYSQL", "mysql.host")

	postgres := os.Getenv("POSTGRES")
	mysql := os.Getenv("MYSQL")

	fmt.Println("postgres config: ", postgres)
	fmt.Println("mysql config: ", mysql)

}
