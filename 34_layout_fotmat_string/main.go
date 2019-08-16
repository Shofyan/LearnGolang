// https://dasarpemrogramangolang.novalagung.com/37-string-format.html

package main

import "fmt"

type student struct {
	name       string
	height     float64
	age        int32
	isGraduate bool
	hobbies    []string
}

var data = student{
	name:       "shofyan",
	height:     178.5,
	age:        27,
	isGraduate: false,
	hobbies:    []string{"read", "coding", "watching movie"},
}

func main() {

	// cetak karakter
	fmt.Printf("%c\n", 1400)
	fmt.Printf("%c\n", 1235)

	// cetak berbasis 10
	fmt.Printf("%d\n", data.age)

	// cetak bilangan scientific
	fmt.Printf("%e\n", data.height)
	fmt.Printf("%E \n", data.height)

	// cetak bilangan float
	fmt.Printf("%f\n", data.height)
	fmt.Printf("%9f\n", data.height)
	fmt.Printf("%.2f\n", data.height)
	fmt.Printf("%.f\n", data.height)

	// cetak bilangan desimal yang besar
	fmt.Printf("%g\n", 0.123123123123)

	// cetak bilangan okta
	fmt.Printf("%o\n", data.age)

	// cetak alamat pointer
	fmt.Printf("%p\n", &data.name)

	// cetak escape string
	fmt.Printf("%q\n", `"name \ height"`)

	// cetak string
	fmt.Printf("%s\n", data.name)

	// cetak boolean
	fmt.Printf("%t\n", data.isGraduate)

	// cetak tipe variable
	fmt.Printf("%T\n", data.name)

	// cetak value dari variable
	fmt.Printf("%v\n", data)

	// cetak struct sesuai dengan urutan
	fmt.Printf("%+v\n", data)

	// cetak struct sesuai dengan urutan struct
	fmt.Printf("%#v\n", data)

	// cetak nilai hexa decimal
	fmt.Printf("%x\n", data.age)
	// 1a

	var d = data.name

	fmt.Printf("%x%x%x%x\n", d[0], d[1], d[2], d[3])
	// 7769636b

	fmt.Printf("%x\n", d)
	// 7769636b

	fmt.Printf("%%\n")
	// %
}
