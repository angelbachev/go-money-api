package create_account

import (
	"github.com/angelbachev/go-money-api/domain/account"
)

type CreateAccountCommandHandler struct {
	accountRepository account.AccountRepositoryInterface
}

func NewCreateAccountCommandHandler(
	accountRepository account.AccountRepositoryInterface,
) *CreateAccountCommandHandler {
	return &CreateAccountCommandHandler{
		accountRepository: accountRepository,
	}
}

func (h CreateAccountCommandHandler) Handle(command CreateAccountCommand) (*CreateAccountResponse, error) {
	account := account.NewAccount(command.UserID, command.Name, command.Description, command.CurrencyCode)
	h.accountRepository.CreateAccount(account)

	return NewCreateAccountResponse(account), nil
}
