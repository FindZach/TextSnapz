package main

import (
	"github.com/FindZach/TextSnapz/internal/application"
	"github.com/FindZach/TextSnapz/internal/infrastructure/db"
	"github.com/FindZach/TextSnapz/internal/presentation/api"
	"github.com/FindZach/TextSnapz/internal/presentation/web"
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

	// Initialize repository and service
	repo := db.NewSqliteTextSnapRepository(sqliteDB)
	service := application.NewTextSnapzService(repo)
	apiHandler := api.NewTextSnapHandler(service)
	webHandler, err := web.NewTextSnapWebHandler(service)
	if err != nil {
		log.Fatalf("Failed to set up web handler: %v", err)
	}

	// Set up Gin router
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK", "message": "TextSnapz is running"})
	})

	// API routes
	apiGroup := r.Group("/api")
	{
		apiGroup.POST("/snaps", apiHandler.CreateSnap)
		apiGroup.GET("/snaps/:id", apiHandler.GetSnap)
	}

	// Web routes
	webHandler.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
