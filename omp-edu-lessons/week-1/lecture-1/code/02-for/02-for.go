package main

import "fmt"

func main() {
	var sum int = 0

	for i := 0; i < 10; i++ {
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
