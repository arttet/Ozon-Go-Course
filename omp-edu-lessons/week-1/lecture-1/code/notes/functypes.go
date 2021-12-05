package main

import "fmt"

var f func(func(int, int) int, int) int
var fRetFunc func(func(int, int) int, int) func(int, int) int

var m map[string]map[int]bool

func main() {
	fmt.Println("Hello")
}
