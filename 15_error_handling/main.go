package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	istr := "10"

	i, err := strconv.Atoi(istr)

	if err != nil {
		fmt.Println("hasil error :", err.Error())
	}

	fmt.Println(i)
	fmt.Println(Div(5, 0))
}

func Div(x, y int) (int, error) {

	if y == 0 {
		return 0, errors.New("tidak bisa mebagi dengan nilai 0")
	}

	result := x / y
	return result, nil

}
