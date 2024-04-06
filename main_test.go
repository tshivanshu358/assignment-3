package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// Create a new HTTP request with a GET method and nil body
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Create a ResponseRecorder to record the response from the handler
	rr := httptest.NewRecorder()

	// Call the handler function with the created request and ResponseRecorder
	handler(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the body of the response
	expectedBody := "Hello, World! This line is to validate te code and the test are working properly.I will commit again"
	if rr.Body.String() != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expectedBody)
	}
}
