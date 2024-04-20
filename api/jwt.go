package api

import (
	"os"

	"github.com/go-chi/jwtauth/v5"
)

var tokenAuth *jwtauth.JWTAuth

func setTokenAuth() {
	secret := os.Getenv("JWT_SECRET")

	tokenAuth = jwtauth.New("HS256", []byte(secret), nil) // replace with secret key
}

// TODO: generate different token every time
func createJWT(userID int64) string {
	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"userId": userID})

	return tokenString
}
