package api

import (
	database_util "backend/db/sqlc"
	"backend/proto/users"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

type Router struct {
	queries *database_util.Queries
}

type Server struct {
	users.UsersServiceServer
	queries *database_util.Queries
}

func RegisterService(grpcServer *grpc.Server, queries *database_util.Queries) *Server {
	users.RegisterUsersServiceServer(grpcServer, &Server{queries: queries})
	return &Server{queries: queries}
}

func (server *Server) Register(ctx context.Context, req *users.RegisterRequest) (*users.TokenResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	fmt.Println(string(hashedPassword))
	if err != nil {
		return nil, err
	}
	result, err := server.queries.CreateUser(ctx, database_util.CreateUserParams{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		return nil, err
	}
	userID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}
	return &users.TokenResponse{Token: token}, nil
}
