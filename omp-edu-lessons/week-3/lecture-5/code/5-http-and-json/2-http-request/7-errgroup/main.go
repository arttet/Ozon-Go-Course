package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func main() {
	httpClient := &http.Client{}

	ctx := context.Background()
	eg, egCtx := errgroup.WithContext(ctx)

	var isExampleComAvailable bool
	eg.Go(func() error {
		reqCtx, cancel := context.WithTimeout(egCtx, 2*time.Second)
		defer cancel()

		req, _ := http.NewRequestWithContext(reqCtx, http.MethodGet,
			"https://example.com/", nil)
		resp, err := httpClient.Do(req)
		if err != nil {
			return errors.Wrap(err, "fetch example.com failed")
		}
		defer resp.Body.Close()

		isExampleComAvailable = resp.StatusCode == http.StatusOK
		return nil
	})

	var isGoogleComAvailable bool
	eg.Go(func() error {
		reqCtx, cancel := context.WithTimeout(egCtx, time.Second)
		defer cancel()

		req, err := http.NewRequestWithContext(reqCtx, http.MethodGet,
			"https://google.com/", nil)
		resp, err := httpClient.Do(req)
		if err != nil {
			return errors.Wrap(err, "fetch google.com failed")
		}
		defer resp.Body.Close()

		isGoogleComAvailable = resp.StatusCode == http.StatusOK
		return nil
	})

	err := eg.Wait()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("isExampleComAvailable", isExampleComAvailable)
	fmt.Println("isGoogleComAvailable", isGoogleComAvailable)
}
