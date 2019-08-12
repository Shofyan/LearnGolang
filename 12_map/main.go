package main

import "fmt"

func main() {

	var mapPerson map[int]string
	mapPerson = make(map[int]string)

	mapPerson[1] = "shofyan"
	mapPerson[2] = "rani"
	mapPerson[3] = "Cerah"

	for k, v := range mapPerson {
		fmt.Println(k, v)
	}

	pp, ok := mapPerson[5]

	if !ok {
		fmt.Println("tidak ada mapPerson")
	}

	fmt.Println(pp)

}
