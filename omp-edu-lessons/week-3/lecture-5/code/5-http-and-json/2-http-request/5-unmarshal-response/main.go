package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type getQuoteContract struct {
	Quote string `json:"quote"`
}

func main() {
	fmt.Println()

	ctx := context.Background()
	contract, err := getQuote(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Kanye West says %q\n", contract.Quote)
}

func getQuote(ctx context.Context) (*getQuoteContract, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet,
		"https://api.kanye.rest/", nil)

	httpClient := &http.Client{Timeout: 5 * time.Second}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("httpClient.Do: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %w", err)
	}

	contract := new(getQuoteContract)
	err = json.Unmarshal(respBody, contract)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return contract, nil
}
