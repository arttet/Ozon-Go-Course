package main

import "fmt"

func main() {
	m := map[string]int{}

	fill(m)

	fmt.Println(m)
}

func fill(m map[string]int) {
	m["one"] = 1
}
