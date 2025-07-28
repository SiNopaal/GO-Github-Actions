package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SiNopaal/go-github-actions/internal"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Root path OK")
	})

	http.HandleFunc("/hello", internal.HelloHandler)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
