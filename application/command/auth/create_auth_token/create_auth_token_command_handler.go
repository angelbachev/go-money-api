package create_auth_token

import (
	"errors"

	"github.com/angelbachev/go-money-api/domain/auth"
	"github.com/angelbachev/go-money-api/domain/user"
)

type CreateAuthTokenCommandHandler struct {
	userRepository         user.UserRepositoryInterface
	userSettingsRepository user.UserSettingsRepositoryInterface
	authService            auth.AuthService
}

func NewCreateAuthTokenCommandHandler(
	userRepository user.UserRepositoryInterface,
	userSettingsRepository user.UserSettingsRepositoryInterface,
	authService auth.AuthService,
) *CreateAuthTokenCommandHandler {
	return &CreateAuthTokenCommandHandler{
		userRepository:         userRepository,
		userSettingsRepository: userSettingsRepository,
		authService:            authService,
	}
}

func (h CreateAuthTokenCommandHandler) Handle(command CreateAuthTokenCommand) (*CreateAuthTokenResponse, error) {
	// TODO: validate email and password

	user, err := h.userRepository.GetUserByEmail(command.Email)
	if err != nil {
		return nil, errors.New("Unable to login user")
	}

	if !user.ValidPassword(command.Password) {
		return nil, errors.New("Unable to login user")
	}

	settings, err := h.userSettingsRepository.GetUserSettings(user.ID)
	if err != nil {
		return nil, errors.New("Unable to login user")
	}

	return NewCreateAuthTokenResponse(h.authService.GenerateToken(user.ID), settings), nil
}
