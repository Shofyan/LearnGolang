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
	fmt.Println(Div(5, 3))

	if x, err := Div(10, 2); err != nil {
		fmt.Println("hasil error :", err.Error())
	} else {
		fmt.Println(x)
	}

}

// Div adalah fungsi  untuk membagi x dengan y
// jika y == 0 maka kembalikan pesan error
func Div(x, y int) (int, error) {

	if y == 0 {
		return 0, errors.New("tidak bisa mebagi dengan nilai 0")
	}

	result := x / y
	return result, nil

}
