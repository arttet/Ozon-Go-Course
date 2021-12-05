package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	fmt.Println()

	{ // Marshal
		data := map[string]interface{}{
			"a": 123,
			"b": "text",
			"c": struct{}{},
			"d": 123.456,
		}

		jsonBytes, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonBytes))
	}

	fmt.Println()

	{ // MarshalIndent
		data := map[string]int{
			"a": 1,
			"b": 2,
		}

		jsonBytes, err := json.MarshalIndent(data, "<prefix>", "<indent>")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonBytes))
	}
}
