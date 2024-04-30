package update_account

import (
	"github.com/angelbachev/go-money-api/domain/account"
)

type UpdateAccountResponse struct {
	Account *account.Account
}

func NewUpdateAccountResponse(account *account.Account) *UpdateAccountResponse {
	return &UpdateAccountResponse{
		Account: account,
	}
}
