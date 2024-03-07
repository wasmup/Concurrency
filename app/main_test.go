package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogsLevel(t *testing.T) {
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	logsLevel(w, r)

	// Check the status code of the response
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decode the response body
	var respBody struct {
		Status int
		Level  int
	}
	if err := json.NewDecoder(w.Body).Decode(&respBody); err != nil {
		t.Fatal(err)
	}

	// Check the response body contains the correct status and level
	if respBody.Status != http.StatusOK {
		t.Errorf("handler returned unexpected status: got %v want %v",
			respBody.Status, http.StatusOK)
	}
	if respBody.Level != 0 { // Assuming default level is 0
		t.Errorf("handler returned unexpected level: got %v want 0", respBody.Level)
	}

	// Create a new request with method PUT and level 1
	reqPut, err := http.NewRequest("PUT", "/", bytes.NewBuffer([]byte(`{"Level": 1}`)))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder to record the response
	rrPut := httptest.NewRecorder()

	// Call the logsLevel function with the PUT request
	logsLevel(rrPut, reqPut)

	// Check the status code of the response
	if status := rrPut.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check that the level was set correctly
	if lvl.Level() != 1 {
		t.Errorf("handler failed to set level: got %v want %v", lvl.Level(), 1)
	}
}
