package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	p := Point{X: 1}

	fmt.Println(p)
	// fmt.Printf("%+v\n", p)
}
