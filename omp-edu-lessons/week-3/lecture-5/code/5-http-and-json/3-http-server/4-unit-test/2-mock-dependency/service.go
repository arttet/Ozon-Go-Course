package main

import "errors"

var ErrProductNotFound = errors.New("product not found")

// Product товар
type Product struct {
	// SKU уникальный идентификатор товара
	SKU int64 `json:"sku"`
	// Name название товара
	Name string `json:"name"`
}
