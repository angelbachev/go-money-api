package create_account

import (
	"github.com/angelbachev/go-money-api/domain/account"
)

type CreateAccountResponse struct {
	Account *account.Account
}

func NewCreateAccountResponse(account *account.Account) *CreateAccountResponse {
	return &CreateAccountResponse{
		Account: account,
	}
}
