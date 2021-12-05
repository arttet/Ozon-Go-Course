package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type myStruct struct {
	privateField string
}

func (s *myStruct) UnmarshalJSON(json []byte) error {
	const prefix, suffix = `{"privateField": "`, `"}`

	json = bytes.TrimPrefix(json, []byte(prefix))
	json = bytes.TrimSuffix(json, []byte(suffix))
	json = bytes.TrimSpace(json)

	s.privateField = string(json)
	return nil
}

func main() {
	fmt.Println()

	var data = new(myStruct)
	err := json.Unmarshal([]byte(`{"privateField": "   test  "}`), data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data.privateField)
}
