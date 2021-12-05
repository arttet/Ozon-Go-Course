package main

import "fmt"

type PointWithComments struct {
	Point    Point
	Comments []string
}

type Point struct {
	X int
	Y int
}

func main() {
	var p PointWithComments

	fmt.Println(p)
}
