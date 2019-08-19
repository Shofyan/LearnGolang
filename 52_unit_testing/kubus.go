package main

import (
	"math"
)

type kubus struct {
	Sisi float64
}

func (k kubus) Volume() float64 {
	return math.Pow(k.Sisi, 3)
}

func (k kubus) Luas() float64 {
	return math.Pow(k.Sisi, 2) * 6
}

func (k kubus) keliling() float64 {
	return k.Sisi * 12
}
