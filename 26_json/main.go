package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	p1 := Player{ID: 1, Name: "shofyan"}

	p1Json, err := json.Marshal(p1)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(p1Json))

	fmt.Println("--------------------------")
	data := []byte(`{"id":1,"name":"shosfyan"}`)
	var p2 Player
	err = json.Unmarshal(data, &p2)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(p2)

}

//Player adalah struct pemain
type Player struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
