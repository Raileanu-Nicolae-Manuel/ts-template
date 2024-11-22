package api

import (
	"backend/config"
	database_util "backend/db/sqlc"
	"backend/websocket"
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
	Get(router, "/ping", Ping, nil)
	config := config.LoadConfig()
	if config.Env == "development" {
		router.HandleFunc("/ws", websocket.HandleWebSocket)
	}
}

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}
