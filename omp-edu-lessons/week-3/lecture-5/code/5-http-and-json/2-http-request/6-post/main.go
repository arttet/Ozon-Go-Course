package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	url         = "https://httpbin.org/post"
	contentType = "application/json"
	reqBody     = `{"id": 999, "value": "content"}`
)

func main() {
	httpClient := &http.Client{Timeout: time.Second}

	{
		req, _ := http.NewRequest(http.MethodPost,
			url, strings.NewReader(reqBody))
		req.Header.Set("Content-Type", contentType)

		resp, err := httpClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		fmt.Println("body", string(body))
	}

	{
		// No context, only GET, POST and HEAD
		resp, err := httpClient.Post(url, contentType,
			strings.NewReader(reqBody))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		fmt.Println("body", string(body))
	}
}
