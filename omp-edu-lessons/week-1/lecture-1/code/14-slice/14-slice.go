package main

import (
	"fmt"
)

func main() {
	fmt.Printf("description | first element pointer | len | cap |   slice  \n")
	printLine()
	a := []int{}
	printSlice("before", a)
	a = append(a, 1)
	printSlice("after 1", a)
	a = append(a, 2)
	printSlice("after 2", a)
	a = append(a, 3)
	printSlice("after 3", a)
	a = append(a, 4)
	printSlice("after 4", a)
	a = append(a, 5)
	printSlice("after 5", a)
}

func printSlice(desc string, slice []int) {
	var pointer *int
	if len(slice) > 0 {
		pointer = &slice[0]
	}
	fmt.Printf(
		"%11s | %21p | %3d | %3d | %v\n",
		desc,
		pointer,
		len(slice),
		cap(slice),
		slice,
	)
	printLine()
}

func printLine() {
	fmt.Println("--------------------------------------------------------------")
}
