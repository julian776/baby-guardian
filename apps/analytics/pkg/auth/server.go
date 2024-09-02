package auth

import (
	"context"
	"fmt"

	pb "github.com/julian776/baby-guardian/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	validEmail    = "admin"
	validPassword = "1234"
)

type AuthServer struct {
	pb.UnimplementedAuthServer
}

func NewAuthServer() *AuthServer {
	return &AuthServer{}
}

func (s *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	fmt.Println("Login request received")
	fmt.Printf("%+v\n", req)
	if req.GetEmail() == "" || req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "email and password are required")
	}

	if req.GetEmail() != validEmail || req.GetPassword() != validPassword {
		fmt.Println("Invalid email or password")
		return nil, ErrUnauthenticated
	}

	m, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("No metadata")
		return nil, ErrUnauthenticated
	}

	if len(m["authorization"]) == 0 {
		return nil, ErrUnauthenticated
	}

	t, err := GenerateToken(req.GetEmail(), req.GetPassword(), authToken)
	if err != nil {
		return nil, ErrUnauthenticated
	}

	return &pb.LoginResponse{
		Token: t,
	}, nil
}
