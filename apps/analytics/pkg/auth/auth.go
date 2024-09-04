package auth

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Role  string `json:"role"`
	Email string `json:"email"`
}

func AuthUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	m, ok := metadata.FromIncomingContext(ctx)
	fmt.Printf("Metadata: %+v\n", m)
	if !ok {
		return nil, ErrUnauthenticated
	}

	if len(m["authorization"]) == 0 {
		return nil, ErrUnauthenticated
	}

	s, found := strings.CutPrefix(m["authorization"][0], "Bearer ")
	if !found {
		// This should be validated with a proper route for
		// login. But for the simplicity of the example
		if s == authToken {
			return handler(ctx, req)
		}
	}
	tkn, err := jwt.ParseWithClaims(s, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		return nil, ErrUnauthenticated
	}

	if !tkn.Valid {
		return nil, ErrUnauthenticated
	}

	return handler(ctx, req)
}

func GenerateToken(
	email string,
	pwd string,
	token string,
) (string, error) {
	if token != authToken {
		return "", ErrUnauthenticated
	}

	// Fetch the user from the database
	if email != "admin" || pwd != "1234" {
		return "", ErrUnauthenticated
	}

	t := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&CustomClaims{
			Role:  "admin",
			Email: email,
			RegisteredClaims: jwt.RegisteredClaims{
				Audience:  jwt.ClaimStrings{"baby-guardian"},
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
				Issuer:    "baby-guardian",
			},
		})

	s, err := t.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return s, nil
}
