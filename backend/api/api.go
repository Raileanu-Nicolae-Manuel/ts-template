package api

import (
	database_util "backend/db/sqlc"
	"fmt"
	"net/http"
)

type Router struct {
	queries *database_util.Queries
	users   *UsersController
}

func CreateRouter(queries *database_util.Queries) *Router {
	return &Router{queries: queries, users: &UsersController{queries: queries}}
}

func (r *Router) RegisterRoutes(router *http.ServeMux) {
	r.users.RegisterRoutes(router, "/users")
	Get(router, "/ping", Ping)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}
