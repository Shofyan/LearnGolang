// https://dasarpemrogramangolang.novalagung.com/50-json.html
package main

import "fmt"
import "encoding/json"

type User struct {
	FullName string `json:"name"`
	Age      int    `json: "age"`
}

func main() {
	var jsonString = `{"name":"shofyan","age":23}`
	var jsonData = []byte(jsonString)
	var data User
	err := json.Unmarshal(jsonData, &data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user : ", data.FullName)
	fmt.Println("age :", data.Age)

	// encode object ke object string
	var object = []User{{"shofyan", 23}, {"superman", 30}}
	jsonData2, err := json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString2 = string(jsonData2)
	fmt.Println(jsonString2)

}
