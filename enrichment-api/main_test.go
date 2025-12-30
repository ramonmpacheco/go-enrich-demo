package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleEnrichment_BadRequest(t *testing.T) {
	// Test missing customer_code
	req, err := http.NewRequest("GET", "/enrichment", nil)
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
