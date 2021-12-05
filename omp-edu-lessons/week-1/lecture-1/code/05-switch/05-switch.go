package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input = readLine("Enter yes or no: ")

	switch input {
	case "yes", "да":
		fmt.Println("You've agreed")
	case "no", "нет":
		fmt.Println("You've disagreed")
	default:
		fmt.Println("I don't understand")
	}
}

func readLine(greeting string) string {
	fmt.Print(greeting)
	reader := bufio.NewReader(os.Stdin)
	text, _, _ := reader.ReadLine()
	return string(text)
}
