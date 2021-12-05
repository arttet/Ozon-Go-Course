package main

import "fmt"

func main() {
	strToNum := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	fmt.Println(strToNum)
	fmt.Println(strToNum["two"])
	// fmt.Println(strToNum["four"])
}
