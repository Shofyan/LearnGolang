package main

import (
	"fmt"
)

func main() {
	var hello string = "Hello"
	var strPtr *string = &hello

	fmt.Println("nilai strPtr: " + *strPtr)
	fmt.Printf("Tipe pointer strPtr:  %T \n", strPtr)
	fmt.Printf("alamat memory dari  strPtr:=&hello  %p \n", strPtr)
	fmt.Printf("alamat memory strPtr:  %p \n", &strPtr)
	changePtr(&hello)
	fmt.Println(hello)
}

func changePtr(v *string) {
	*v += " Golang"
}

func change(v string) {
	v += " Golang"
}
