package main

import (
	"backend/internal/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	fmt.Println("Hello, WELCOME TO MY APPLICATOPN")
	r := mux.NewRouter()

	r.HandleFunc("/help", handlers.HandleHelpCommand).Methods("GET")
	r.HandleFunc("/about", handlers.HandleAboutCommand).Methods("GET")
	r.HandleFunc("/draw", handlers.HandleDrawCommand).Methods("GET")
	r.HandleFunc("/fetch-price/{pair}", handlers.HandleFetchPriceCommand).Methods("GET")
	r.HandleFunc("/upload", handlers.HandleUploadCommand).Methods("POST")
	r.HandleFunc("/delete/{fileName}", handlers.HandleDeleteFile).Methods("DELETE")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // You can specify allowed origins
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(r)

	http.ListenAndServe(":8080", handler)
}
