package main

import "fmt"

func main() {
	strings := []string{"one", "two", "three"}

	for key, value := range strings {
		fmt.Printf(`strings[%d] = "%s"`+"\n", key, value)
	}
}
