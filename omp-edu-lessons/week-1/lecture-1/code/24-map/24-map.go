package main

import "fmt"

func main() {
	var m map[string]int // = make(map[string]int)

	m["one"] = 1

	fmt.Println(m)
}
