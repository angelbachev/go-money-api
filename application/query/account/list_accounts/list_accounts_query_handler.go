package list_accounts

import (
	"github.com/angelbachev/go-money-api/domain/account"
)

type ListAccountsQueryHandler struct {
	accountRepository account.AccountRepository
}

func NewListAccountsQueryHandler(
	accountRepository account.AccountRepository,
) *ListAccountsQueryHandler {
	return &ListAccountsQueryHandler{
		accountRepository: accountRepository,
	}
}

func (h ListAccountsQueryHandler) Handle(query ListAccountsQuery) (*ListAccountsResponse, error) {
	accounts, err := h.accountRepository.GetAccounts(query.UserID)
	if err != nil {
		return nil, err
	}

	return NewListAccountsResponse(accounts), nil
}
