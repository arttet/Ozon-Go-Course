package main

import "fmt"

func main() {
	strToNum := make(map[string]int)

	fmt.Println("len(strToNum) =", len(strToNum))

	strToNum["zero"] = 0
	strToNum["one"] = 1
	strToNum["two"] = 2
	strToNum["three"] = 3

	fmt.Println("len(strToNum) =", len(strToNum))
}
