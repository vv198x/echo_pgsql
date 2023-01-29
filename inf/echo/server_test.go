package echo

import (
	"net/http"
	"testing"
)

func TestStart(t *testing.T) {
	// Start the server in a goroutine
	go Start(":8000")

	// Make a GET request to the server
	res, err := http.Get("http://localhost:8000/api/users/v1/")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	// Check the response status code
	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}
}
