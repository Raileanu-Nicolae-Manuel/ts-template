package api

import (
	database_util "backend/db/sqlc"
	"context"
	"encoding/json"
	"net/http"
)

type UsersController struct {
	queries *database_util.Queries
}

func (c *UsersController) RegisterRoutes(router *http.ServeMux, path string) {
	// here we are registering the routes for the users controller
	Get(router, path, c.GetUsers)
}

func (c *UsersController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.queries.ListUsers(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}
