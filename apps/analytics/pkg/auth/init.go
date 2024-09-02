package auth

import "os"

var (
	authToken = ""
	secretKey = ""
)

func init() {
	authToken = os.Getenv("AUTH_TOKEN")
	if authToken == "" {
		panic("AUTH_TOKEN not set")
	}

	secretKey = os.Getenv("SECRET")
	if secretKey == "" {
		panic("SECRET not set")
	}
}
