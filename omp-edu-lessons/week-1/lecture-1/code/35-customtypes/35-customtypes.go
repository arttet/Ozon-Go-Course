package main

import "fmt"

type EnrichedInt int

func (i EnrichedInt) IsNegative() bool {
	return i < 0
}

func (i *EnrichedInt) Set(newValue int) {
	*i = EnrichedInt(newValue)
}

func main() {
	var i EnrichedInt = 4

	// i -= 10

	fmt.Println(i.IsNegative())
	i.Set(-1)
	fmt.Println(i.IsNegative())
}
