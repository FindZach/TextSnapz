package main

import (
	"github.com/FindZach/TextSnapz/internal/infrastructure/db"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Initialize SQLite database
	sqliteDB, err := db.SetupDatabase("./textsnapz.db")
	if err != nil {
		log.Fatalf("Failed to set up database: %v", err)
	}
	defer sqliteDB.Close()

	// Set up Gin router
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK", "message": "TextSnapz is running"})
	})

	// TODO: Add API routes for snaps

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
