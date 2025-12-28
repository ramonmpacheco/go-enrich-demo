package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleEnrichment_BadRequest(t *testing.T) {
	req, err := http.NewRequest("POST", "/enrichment", strings.NewReader("invalid-json"))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleEnrichment)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
