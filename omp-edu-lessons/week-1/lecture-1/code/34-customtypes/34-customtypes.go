package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func (p Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (p Point) SetX(newX float64) {
	p.X = newX
}

func main() {
	p := Point{1, 2}

	p.SetX(100)

	fmt.Println(p.Abs())
}
