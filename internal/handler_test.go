package internal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HelloHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode JSON response
	var data map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &data)
	if err != nil {
		t.Fatalf("could not parse JSON response: %v", err)
	}

	expected := "Hello, World!"
	if data["message"] != expected {
		t.Errorf("expected message %q, got %q", expected, data["message"])
	}
}
