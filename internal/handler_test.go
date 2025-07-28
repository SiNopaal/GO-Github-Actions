package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()
	HelloHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
	expected := `{"message":"Hello, World!"}`
	if w.Body.String() != expected+"\n" {
		t.Errorf("expected body %s, got %s", expected, w.Body.String())
	}
}