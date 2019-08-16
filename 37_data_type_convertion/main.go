// ref https://dasarpemrogramangolang.novalagung.com/40-data-type-conversion.html
package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var str = "123"

	// string to number
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(num)

	// integer to string
	num1 := 123
	str1 := strconv.Itoa(num1)
	fmt.Println(str1)

	// konversi string ke int dengan panjang tertentu dan tipe data int tertentu
	str2 := "456"
	num2, err := strconv.ParseInt(str2, 10, 32)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(num2)

	//string to float conversion
	num3 := "345.12"
	str3, err := strconv.ParseFloat(num3, 32)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(str3)

	//  float to string
	num4 := float64(12.32)
	str4 := strconv.FormatFloat(num4, 'f', 4, 64)
	fmt.Println(str4)

	//	bool to string
	str5 := false
	var b1 = strconv.FormatBool(str5)
	fmt.Println(b1)

	// using casting for conversion
	var a float64 = float64(34)
	fmt.Println(a)

	var b int = int(2)
	fmt.Println(b)

	// string to byte
	hello := "hello dunia"
	bb := []byte(hello)

	for _, v := range bb {
		fmt.Println(v)
		fmt.Println(string(v))
	}

	//  assertion  to data type interface

	data := map[string]interface{}{
		"name":    "shofyan",
		"grade":   2,
		"height":  178.2,
		"isMale":  true,
		"hobbies": []string{"baca", "renang"},
	}
	fmt.Println(data["name"].(string))
	fmt.Println(data["grade"].(int))
	fmt.Println(data["height"].(float64))
	fmt.Println(data["isMale"].(bool))
	fmt.Println(data["hobbies"].([]string))

	for _, val := range data {
		switch val.(type) {
		case string:
			fmt.Println(val.(string))
		case int:
			fmt.Println(val.(int))
		case float64:
			fmt.Println(val.(float64))
		case bool:
			fmt.Println(val.(bool))
		case []string:
			fmt.Println(val.([]string))
		default:
			fmt.Println(val.(int))
		}
	}

}
