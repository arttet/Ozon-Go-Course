package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	p := Point{}

	fill(p)

	fmt.Println(p)
}

func fill(p Point) {
	p.X = 1
	p.Y = 2
}
