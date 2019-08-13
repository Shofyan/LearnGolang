package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	fmt.Println(t)
	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Day())

	fmt.Println("=============")

	t1, err := time.Parse("2019 20 01", "2019 20 01")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(t1)

}
