package create_user

type CreateUserCommand struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
