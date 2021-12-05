package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	fmt.Println("before zerofy", s)

	zerofy(s)

	fmt.Println("after zerofy", s)
}

func zerofy(s []int) {
	for i := range s {
		s[i] = 0
	}
}

// - append slice
// - array
