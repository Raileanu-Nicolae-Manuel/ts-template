package main

import (
	"backend/config"
	"backend/proto/calculator"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

type Server struct {
	calculator.CalculatorServiceServer
}

func (server *Server) Sum(ctx context.Context, req *calculator.SumRequest) (*calculator.SumResponse, error) {
	return &calculator.SumResponse{Result: req.FirstNumber + req.SecondNumber}, nil
}

func main() {
	cfg := config.LoadConfig()

	grpcServer := grpc.NewServer()
	calculator.RegisterCalculatorServiceServer(grpcServer, &Server{})

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
		Addr:    ":8080",
		Handler: handler,
	}
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to serve HTTP: %v", err)
	}
	fmt.Println("Server is running on", cfg.ServerAddress)
}
