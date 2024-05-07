package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tu-usuario/holidays-api/handlers"
)

func main() {
	r := gin.Default()

	// Handlers
	r.GET("/holidays", handlers.GetHolidays)

	// Run server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
