package main

import (
	"backend/api"
	"backend/config"
	"backend/db"
	database_util "backend/db/sqlc"
	"fmt"
	"net/http"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database connection
	database := db.ConnectDB(cfg)
	defer database.Close()
	queries := database_util.New(database)

	router := http.NewServeMux()

	api.CreateRouter(queries).RegisterRoutes(router)

	stack := api.Chain(
		api.Logger,
	)

	server := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: stack(router),
	}
	fmt.Println("Server is running on", cfg.ServerAddress)
	server.ListenAndServe()
}
