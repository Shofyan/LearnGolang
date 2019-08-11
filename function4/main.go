package main

import (
	"fmt"
)

func main() {

	p := func(s string) bool {
		if s == "shofyan" {
			return true
		}

		return false

	}

	fmt.Println(match("rani", p))

}

func match(v string, f func(string) bool) bool {
	if f(v) {
		return true
	}
	return false
}
