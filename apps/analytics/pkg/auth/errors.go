package auth

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUnauthenticated = status.Error(codes.Unauthenticated, "unauthenticated")
)
