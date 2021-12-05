package main

import (
	"net/http"
)

func main() {
	// Так не подходит, нужен указатель
	// var _ http.Handler = http.ServeMux{}

	// Можно так
	mux := http.NewServeMux()
	// Или так
	mux = &http.ServeMux{} // Все поля приватные ¯\_(ツ)_/¯
	// Ну или так
	mux = new(http.ServeMux)

	_ = mux
}
