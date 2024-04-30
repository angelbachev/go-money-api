package create_auth_token

type CreateAuthTokenCommand struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
