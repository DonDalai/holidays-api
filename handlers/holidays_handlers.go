package handlers

import (
	"encoding/json"
	"log"
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

	// Parse holidaysData into models.HolidayResponse
	var holidayResponse models.HolidayResponse
	if err := json.Unmarshal(holidaysData, &holidayResponse); err != nil {
		log.Printf("Failed to unmarshal holidays data: %v", holidayResponse)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse holidays data"})
		return
	}
	// Apply filters based on query parameters
	filteredHolidays := applyFilters(holidayResponse.Data, c.Query("type"), c.Query("start_date"), c.Query("end_date"))

	// Check if any holidays match the filters
	if len(filteredHolidays) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No holidays found matching the specified criteria"})
		return
	}
	// Respond with filtered holidays
	c.JSON(http.StatusOK, filteredHolidays)
}

// Helper function to apply filters
func applyFilters(holidays []models.Holiday, holidayType string, startDateStr string, endDateStr string) []models.Holiday {
	var filteredHolidays []models.Holiday
	// Parse start_date and end_date strings into time.Time objects
	var startDate, endDate time.Time
	var err error
	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			log.Printf("Error parsing start date: %v", err)
			return filteredHolidays
		}
	}
	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			log.Printf("Error parsing end date: %v", err)
			return filteredHolidays
		}
	}

	// Filter holidays based on type and date range
	for _, holiday := range holidays {
		// Check if the holiday matches the specified type and date range
		if (holidayType == "" || holiday.Type == holidayType) &&
			(startDateStr == "" || holidayDateInRange(holiday.Date, startDate, endDate)) {
			// Add the holiday to the filtered list
			filteredHolidays = append(filteredHolidays, holiday)

		}
	}

	return filteredHolidays
}

// Helper function to check if a holiday date is within the specified range
func holidayDateInRange(holidayDate string, startDate time.Time, endDate time.Time) bool {
	// Parse the holiday date string into a time.Time object
	date, err := time.Parse("2006-01-02", holidayDate)
	if err != nil {
		log.Printf("Error parsing holiday date: %v", err)
		return false
	}

	// Check if the holiday date is within the specified range
	return !date.Before(startDate) && !date.After(endDate)
}
