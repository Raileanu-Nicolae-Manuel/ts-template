package main

import (
	"backend/config"
	"backend/db"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database connection
	database := db.ConnectDB(cfg)
	defer database.Close()
}
