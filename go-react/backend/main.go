package main

import (
	"backend/config"
	"backend/db"
	"fmt"
	"log"
	"net/http"

	"backend/api"
	database_util "backend/db/sqlc"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadConfig()
	database := db.ConnectDB(cfg)
	defer database.Close()
	queries := database_util.New(database)

	grpcServer := grpc.NewServer()
	api.RegisterService(grpcServer, queries)

	wrappedGrpc := grpcweb.WrapServer(grpcServer,
		grpcweb.WithOriginFunc(func(origin string) bool {
			return true // Be careful with this in production
		}),
		grpcweb.WithAllowedRequestHeaders([]string{"*"}),
	)

	handler := http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		resp.Header().Set("Access-Control-Allow-Headers", "*")

		if wrappedGrpc.IsGrpcWebRequest(req) {
			wrappedGrpc.ServeHTTP(resp, req)
			return
		}
		// Handle regular HTTP requests if needed
	})

	httpServer := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: handler,
	}
	fmt.Println("Server is running on", cfg.ServerAddress)
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to serve HTTP: %v", err)
	}
}
