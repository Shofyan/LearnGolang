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

	t1, err := time.Parse("2006-01-02T15:04:05Z07:00", "2019-01-20T11:45:26.371Z")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(t1)

}
