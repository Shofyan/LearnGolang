package main

import "fmt"

func main() {
	var players []Player
	players = []Player{
		Player{
			ID:   1,
			Name: "shofyan",
		},
		Player{
			ID:   2,
			Name: "Adi",
		},
	}

	players = append(players, Player{ID: 3, Name: "rani"})

	for _, v := range players {
		fmt.Println(v.ID, v.Name)
	}

}

type Player struct {
	ID   int
	Name string
}
