package main

import (
	"backend/internal/handlers"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, WELCOME TO MY APPLICATOPN")
	http.HandleFunc("/help", handlers.HandleHelpCommand)
	http.HandleFunc("/about", handlers.HandleAboutCommand)
	http.HandleFunc("/upload", handlers.HandleUploadCommand)
	http.HandleFunc("/fetch-price/", handlers.HandleFetchPriceCommand)
	http.ListenAndServe(":8080", nil)
}
