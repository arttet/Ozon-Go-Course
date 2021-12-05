package main

import "fmt"

func main() {
	var sum int = 0

	var i int = 0
	for i >= 10 {
		if i%2 == 1 {
			fmt.Println(i, "is odd")
		} else {
			fmt.Println(i, "is even")
		}

		sum += i
		i++
	}

	fmt.Println("sum is", sum)
}
