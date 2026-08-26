package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	r := httptest.NewRecorder()
	Server{}.health(r, httptest.NewRequest(http.MethodGet, "/health", nil))
	if r.Code != 204 {
		t.Fatal(r.Code)
	}
}
