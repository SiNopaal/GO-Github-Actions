package main

import (
	"log"
	"net/http"
	"os"

	"github.com/SiNopaal/go-github-actions/internal"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/hello", internal.HelloHandler)
	log.Printf("Server running at http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
