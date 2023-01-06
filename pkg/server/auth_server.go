package server

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	rental "rental_easy.in/m/pkg/rentalmgmt"
)

// AuthServer is the server for authentication
type AuthServer struct {
	rental.UnimplementedAuthServiceServer
	jwtManager *JWTManager
}

// NewAuthServer returns a new auth server
func NewAuthServer(jwtManager *JWTManager) rental.AuthServiceServer {
	return &AuthServer{jwtManager: jwtManager}
}

func (server *AuthServer) Login(ctx context.Context, req *rental.LoginRequest) (*rental.LoginResponse, error) {
	log.Println("Server : Logging in the User")
	user := User{
		Username:       req.GetUsername(),
		HashedPassword: req.GetPassword(),
	}
	token, err := server.jwtManager.Generate(&user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &rental.LoginResponse{AccessToken: token}
	return res, nil
}
