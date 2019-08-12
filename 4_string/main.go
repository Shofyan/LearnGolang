package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var myString = "Hello Go"
	var myStringNumber = "10"

	// ubah string ke number
	resultConvStr, err := strconv.Atoi(myStringNumber)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultConvStr * 5)

	// ubah number ke string
	fmt.Println(strconv.Itoa(resultConvStr))

	fmt.Println(myString)
	for i := 0; i < len(myString); i++ {
		println(i)
	}

	fmt.Println(strings.ToUpper(myString))
	fmt.Println(strings.ToLower(myString))
	fmt.Println(strings.Contains(myString, "Go"))

	resultSplit := strings.Split(myString, " ")
	for _, v := range resultSplit {
		fmt.Println(v)
	}

}
