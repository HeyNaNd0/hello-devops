package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoot(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	handler(rr, req)
	if rr.Code != 200 || rr.Body.String() != "Hello, DevOps!" {
		t.Fatalf("unexpected: code=%d, body=%q", rr.Code, rr.Body.String())
	}
}
