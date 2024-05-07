package services

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchHolidays(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the response body
		w.Write([]byte("test response"))
	}))
	defer server.Close()

	// Create a new instance of HolidayService with the test server URL
	hs := NewHolidayService(server.URL)

	// Call the FetchHolidays method
	body, err := hs.FetchHolidays()

	// Check for any errors
	if err != nil {
		t.Errorf("FetchHolidays returned an error: %v", err)
	}

	// Check the response body
	expectedBody := []byte("test response")
	if string(body) != string(expectedBody) {
		t.Errorf("FetchHolidays returned unexpected body. Expected: %s, Got: %s", expectedBody, body)
	}
}
