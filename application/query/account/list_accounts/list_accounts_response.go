package list_accounts

import (
	"github.com/angelbachev/go-money-api/domain/account"
)

type ListAccountsResponse struct {
	Accounts []*account.Account
}

func NewListAccountsResponse(accounts []*account.Account) *ListAccountsResponse {
	return &ListAccountsResponse{
		Accounts: accounts,
	}
}
