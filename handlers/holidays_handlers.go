package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tu-usuario/holidays-api/models"
	"github.com/tu-usuario/holidays-api/services"
)

func GetHolidays(c *gin.Context) {
	// Service initialization (assume API endpoint provided)
	holidayService := services.NewHolidayService("https://api.victorsanmartin.com/feriados/en.json")

	// Fetch holidays
	holidaysData, err := holidayService.FetchHolidays()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch holidays"})
		return
	}

	// Parse holidaysData into models.Holiday array
	var holidays []models.Holiday
	if err := json.Unmarshal(holidaysData, &holidays); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse holidays data"})
		return
	}

	// Apply filters based on query parameters
	filteredHolidays := applyFilters(holidays, c.Query("type"), c.Query("start_date"), c.Query("end_date"))

	// Respond with filtered holidays
	c.JSON(http.StatusOK, filteredHolidays)
}

// Helper function to apply filters
func applyFilters(holidays []models.Holiday, holidayType string, startDateStr string, endDateStr string) []models.Holiday {
	var filteredHolidays []models.Holiday

	// Parse start_date and end_date strings into time.Time objects
	var startDate, endDate time.Time
	if startDateStr != "" {
		startDate, _ = time.Parse("2006-01-02", startDateStr)
	}
	if endDateStr != "" {
		endDate, _ = time.Parse("2006-01-02", endDateStr)
	}

	// Filter holidays based on type and date range
	for _, holiday := range holidays {
		if (holidayType == "" || holiday.Type == holidayType) &&
			(startDateStr == "" || holidayDateInRange(holiday.Date, startDate, endDate)) {
			filteredHolidays = append(filteredHolidays, holiday)
		}
	}

	return filteredHolidays
}

// Helper function to check if a holiday date is within the specified range
func holidayDateInRange(holidayDate string, startDate time.Time, endDate time.Time) bool {
	date, err := time.Parse("2006-01-02", holidayDate)
	if err != nil {
		return false
	}
	return !date.Before(startDate) && !date.After(endDate)
}
