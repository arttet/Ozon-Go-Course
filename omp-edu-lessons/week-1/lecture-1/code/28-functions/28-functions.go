package main

import "fmt"

func main() {
	fmt.Println(sum(2, 3, 4))
}

func sum(items ...int) int {
	result := 0
	for _, value := range items {
		result += value
	}
	return result
}

// items to mention:
// - mutiple arguments
// - named arguments
