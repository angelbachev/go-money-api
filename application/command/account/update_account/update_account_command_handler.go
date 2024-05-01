package update_account

import (
	"github.com/angelbachev/go-money-api/domain/account"
)

type UpdateAccountCommandHandler struct {
	accountRepository account.AccountRepository
}

func NewUpdateAccountCommandHandler(
	accountRepository account.AccountRepository,
) *UpdateAccountCommandHandler {
	return &UpdateAccountCommandHandler{
		accountRepository: accountRepository,
	}
}

func (h UpdateAccountCommandHandler) Handle(command UpdateAccountCommand) (*UpdateAccountResponse, error) {
	account := account.NewAccount(command.UserID, command.Name, command.Description, command.CurrencyCode)
	h.accountRepository.UpdateAccount(account)

	return NewUpdateAccountResponse(account), nil
}
