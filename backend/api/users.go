package api

import (
	database_util "backend/db/sqlc"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type UsersController struct {
	queries *database_util.Queries
}

func (c *UsersController) RegisterRoutes(router *http.ServeMux, path string) {
	// here we are registering the routes for the users controller
	usersResponse := Get(router, path, c.GetUsers, []database_util.User{})
	userResponse := Get(router, path+"/{id}", c.GetUser, database_util.User{})

	// Now you can access the Type field like this:
	responseType := usersResponse.Type  // This will be []database_util.User{}
	singleUserType := userResponse.Type // This will be database_util.User{}
	fmt.Println(responseType)
	fmt.Println(singleUserType)
}

func (c *UsersController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.queries.ListUsers(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func (c *UsersController) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	user, err := c.queries.GetUser(context.Background(), int64(idInt))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (c *UsersController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user database_util.CreateUserParams
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdUserResult, err := c.queries.CreateUser(context.Background(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdUserResult)
}

// func (c *UsersController) UpdateUser(w http.ResponseWriter, r *http.Request) {
