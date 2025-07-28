package internal

import (
	"encoding/json"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Run() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8080", nil)
}
