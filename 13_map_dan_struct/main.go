package main

import "fmt"

func main() {

	mapA := make(map[int]Player)

	mapA[1] = Player{ID: 1, Name: "shofyan"}
	mapA[2] = Player{ID: 2, Name: "ani"}

	for _, v := range mapA {
		fmt.Println(v.Name)
	}

}

type Player struct {
	ID   int
	Name string
}
