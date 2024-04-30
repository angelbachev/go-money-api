package auth

type AuthService interface {
	GenerateToken(userID int64) string
}
