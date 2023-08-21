package router

import (
	"io"
	"net/http"
	"testing"
	"time"
)

func TestInitialize(t *testing.T) {
	// Start the router in a goroutine

	// Wait for the server to start up
	time.Sleep(100 * time.Millisecond)

	// Send an HTTP GET request to the /api/v0/openings endpoint
	resp, err := http.Get("http://localhost:8080/api/v0/openings")
	if err != nil {
		t.Fatal(err)
	}

	// Close response listening
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	// Verify that the response status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}
}
