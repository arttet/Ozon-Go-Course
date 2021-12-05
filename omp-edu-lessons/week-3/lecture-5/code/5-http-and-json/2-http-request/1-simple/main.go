package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("https://example.com/")
	if err != nil {
		panic(err)
	}

	fmt.Println("StatusCode", res.StatusCode)
}
