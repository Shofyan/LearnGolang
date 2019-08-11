package main

import (
	"fmt"
)

func main() {
	nextValue := genNumber()

	for i := 0; i < 10; i++ {
		fmt.Println(nextValue())
	}

	lf := love("shofyan")
	fmt.Println(lf("java"))
	
}

func genNumber() func() int {
	x := 0
	return func() int {
		x = x + 1
		return x

	}
}

func love(name string) func(string) string {
	return func(thing string) string {
		return fmt.Sprintf("%s love %s", name, thing)

	}
}
