package echo

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"userSL/inf/pgsql"

	"github.com/labstack/echo"
)

func TestGetToken(t *testing.T) {
	// Create a request to pass to the handler
	req, err := http.NewRequest("POST", "/api/users/v1/auth/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock database for testing
	db := pgsql.GetTestDb()

	// Create a mock JSONLogin struct for testing
	login := JSONLogin{
		Login:    "testuser",
		Password: "testpassword",
	}

	// Bind the mock JSONLogin struct to the request
	c := echo.New().NewContext(req, httptest.NewRecorder())
	c.Set("db", db)
	if err := c.Bind(login); err != nil {
		t.Error(err)
	}

	// Call the getToken function with the mock request and context
	if err := getToken(c); err != nil {
		t.Error(err)
	}

	// Check the response status code
	if c.Response().Status != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, c.Response().Status)
	}

	// Check the response body
	expectedBody := `{"token":"eyJhbGciOiAiSFMyNTYiLCAidHlwIjogIkpXVCJ9.eyJsb2dpbiI6ICJ0ZXN0dXNlciIsICJydWxlIjogMCwgImlhdCI6IDE1NzM5NTQwMDB9.28ce5d23ce74c5f5b66e1c049d11916a"}`
	if c.Response() != nil {
		t.Errorf("Expected response body %s", expectedBody)
	}
}
