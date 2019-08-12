package main

import (
	"fmt"
)

func main() {
	var arr [10]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	fmt.Println(arr[3])

	arrStr := [5]string{"shofyan", "alex", "zulaiha", "adi", "koko"}
	fmt.Println(arrStr[2])

	arrMulti := [2][3]int{{2, 3}, {3, 4, 5}}

	fmt.Println(arrMulti[0][1])
}
