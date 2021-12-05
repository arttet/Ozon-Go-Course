package main

import "fmt"

func main() {
	array := [5]float64{1, 2, 3, 4, 5}

	slice := array[0:2] // [1:3], [:], [1:], [:2]

	fmt.Println(slice)
	fmt.Println("len(slice)", len(slice))
	fmt.Println("cap(slice)", cap(slice))
}
