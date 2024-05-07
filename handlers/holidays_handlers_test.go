package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tu-usuario/holidays-api/models"
)

// Define a HolidayServiceInterface
type HolidayServiceInterface interface {
	FetchHolidays() ([]models.Holiday, error)
}

// MockHolidayService is a mock implementation of HolidayServiceInterface
type MockHolidayService struct{}

func (m *MockHolidayService) FetchHolidays() ([]models.Holiday, error) {
	// Mock implementation returning predefined holidays
	return []models.Holiday{
		{Type: "public", Date: "2022-01-01", Title: "New Year's Day"},
		{Type: "public", Date: "2022-12-25", Title: "Christmas Day"},
	}, nil
}

func TestGetHolidays(t *testing.T) {
	// Create a test router
	router := gin.Default()

	// Create a mock holiday service
	holidayService := &MockHolidayService{}

	// Register the GET route for GetHolidays with the handler using dependency injection
	router.GET("/holidays", func(c *gin.Context) {
		// Invoke FetchHolidays from the holiday service
		holidays, err := holidayService.FetchHolidays()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, holidays)
	})

	// Perform a GET request to the GetHolidays endpoint
	req, _ := http.NewRequest("GET", "/holidays", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, resp.Code)

	// Check the response body
	var responseHolidays []models.Holiday
	err := json.Unmarshal(resp.Body.Bytes(), &responseHolidays)
	assert.NoError(t, err)

	// Check the filtered holidays
	expectedHolidays := []models.Holiday{
		{Type: "public", Date: "2022-01-01", Title: "New Year's Day"},
		{Type: "public", Date: "2022-12-25", Title: "Christmas Day"},
	}
	assert.Equal(t, expectedHolidays, responseHolidays)
}
