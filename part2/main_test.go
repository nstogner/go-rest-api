package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nstogner/go-rest-api/part2/rest"
)

func TestGetSingleUser(t *testing.T) {
	id := 1

	// Create the http request
	req, err := http.NewRequest("GET", fmt.Sprintf("http://example.com/users/%v", id), nil)
	if err != nil {
		t.Fatal("unable to generate request", err)
	}

	// Send the request to the API
	rec := httptest.NewRecorder()
	rest.API().ServeHTTP(rec, req)

	// Check the status code
	if exp := http.StatusOK; rec.Code != exp {
		t.Fatalf("expected status code %v, got: %v", exp, rec.Code)
	}

	// Unmarshal and check the response body
	var u rest.User
	if err := json.NewDecoder(rec.Body).Decode(&u); err != nil {
		t.Fatalf("unable to decode response: %s", err)
	}
	if u.ID != id {
		t.Fatalf("unexpected user")
	}
}
