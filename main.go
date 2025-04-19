package main

import (
	"go-gpt/config"
	"go-gpt/handler"
	"log"
	"net/http"
)

func main() {
	// http.HandleFunc("/chat", handler.ChatHandler)
	// log.Println("Server running on :8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))
	config.LoadConfig()

	// Handle API endpoint
	http.HandleFunc("/chat", handler.ChatHandler)

	// Serve static files (index.html, CSS, JS if needed)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
