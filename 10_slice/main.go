package main

import "fmt"

func main() {

	mySlice := []int{10, 11, 2, 3, 4}

	mySliceString := []string{"shofyan"}
	mySliceString = append(mySliceString, "ani", "dangdut")

	for i, v := range mySliceString {
		fmt.Println(i)
		fmt.Println(v)
	}

	for i, v := range mySlice {
		fmt.Println(i)
		fmt.Println(v)
	}
}
