package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	X      float64
	Y      float64
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Retangle struct {
	Width  float64
	Height float64
}

func (r Retangle) Area() float64 {
	return r.Width * r.Height
}

func getArea(s Shape) float64 {
	return s.Area()
}

func main() {
	circle := Circle{X: 1.0, Y: 1.0, Radius: 10.5}
	retangle := Retangle{Width: 5, Height: 10}

	fmt.Println(getArea(circle))
	fmt.Println(getArea(retangle))
}
