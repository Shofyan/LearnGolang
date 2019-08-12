package main

import (
	"fmt"
	"strconv"
)

func main() {
	istr := "10"

	i, err := strconv.Atoi(istr)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(i)

}
