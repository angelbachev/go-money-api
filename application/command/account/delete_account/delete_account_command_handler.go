package delete_account

import (
	"errors"

	"github.com/angelbachev/go-money-api/domain/account"
	"github.com/angelbachev/go-money-api/domain/category"
	"github.com/angelbachev/go-money-api/domain/expense"
)

type DeleteAccountCommandHandler struct {
	accountRepository  account.AccountRepositoryInterface
	categoryRepository category.CategoryRepositoryInterface
	expenseRepository  expense.ExpenseRepositoryInterface
}

func NewDeleteAccountCommandHandler(
	accountRepository account.AccountRepositoryInterface,
	categoryRepository category.CategoryRepositoryInterface,
	expenseRepository expense.ExpenseRepositoryInterface,
) *DeleteAccountCommandHandler {
	return &DeleteAccountCommandHandler{
		accountRepository:  accountRepository,
		categoryRepository: categoryRepository,
		expenseRepository:  expenseRepository,
	}
}

func (h DeleteAccountCommandHandler) Handle(command DeleteAccountCommand) error {
	categoryIDs, err := h.categoryRepository.GetCategories(command.ID)
	if err != nil {
		return err
	}

	// TODO: handle subcategories

	filters := &expense.ExpenseFilters{
		CategoryIDs: categoryIDs,
	}

	// TODO: validate that the user owns the account and category

	expenses, err := h.expenseRepository.GetExpenses(command.UserID, command.ID, filters, 0, 0)
	if err != nil {
		return err
	}

	if !command.Force && (len(categoryIDs) > 1 || len(expenses) > 0) {
		return errors.New("Account is not empty")
	}

	for _, exp := range expenses {
		h.expenseRepository.DeleteExpense(exp.ID)
	}

	for _, cat := range categoryIDs {
		h.categoryRepository.DeleteCategory(cat)
	}

	h.accountRepository.DeleteAccount(command.ID)

	return nil
}
