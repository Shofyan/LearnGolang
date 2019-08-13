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

	t1, err := time.Parse("2006-01-02", "2019-01-20")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(t1)

}
