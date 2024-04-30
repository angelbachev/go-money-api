package auth

import (
	"github.com/go-chi/jwtauth/v5"
)

type JWTAuthService struct {
	secret    string
	tokenAuth *jwtauth.JWTAuth
}

func NewJWTAuth(secret string) *JWTAuthService {
	return &JWTAuthService{
		secret:    secret,
		tokenAuth: jwtauth.New("HS256", []byte(secret), nil), // replace with secret key
	}
}

// TODO: generate different token every time
func (s JWTAuthService) GenerateToken(userID int64) string {
	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := s.tokenAuth.Encode(map[string]interface{}{"userId": userID})

	return tokenString
}

func (s JWTAuthService) GetJWTAuth() *jwtauth.JWTAuth {
	return s.tokenAuth
}
