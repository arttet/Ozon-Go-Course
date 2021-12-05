package main

import "fmt"

func main() {
	strToNum := make(map[string]int)

	strToNum["zero"] = 0
	strToNum["one"] = 1
	strToNum["two"] = 2
	strToNum["three"] = 3
	strToNum["four"] = 4
	strToNum["five"] = 5

	for key, value := range strToNum { // key, value can be omited
		fmt.Printf(`strToNum["%s"] = %d`+"\n", key, value)
	}
}
