package delete_category

import (
	"errors"

	"github.com/angelbachev/go-money-api/domain/category"
	"github.com/angelbachev/go-money-api/domain/expense"
)

type DeleteCategoryCommandHandler struct {
	categoryRepository category.CategoryRepository
	expenseRepository  expense.ExpenseRepository
}

func NewDeleteCategoryCommandHandler(
	categoryRepository category.CategoryRepository,
	expenseRepository expense.ExpenseRepository,
) *DeleteCategoryCommandHandler {
	return &DeleteCategoryCommandHandler{
		categoryRepository: categoryRepository,
		expenseRepository:  expenseRepository,
	}
}

func (h DeleteCategoryCommandHandler) Handle(command DeleteCategoryCommand) error {
	categoryIDs, err := h.categoryRepository.GetListCategoryIDsAndTheirSubcategories([]int64{command.ID})
	if err != nil {
		return err
	}

	// TODO: handle subcategories

	filters := &expense.ExpenseFilters{
		CategoryIDs: categoryIDs,
	}

	// TODO: validate that the user owns the account and category

	expenses, err := h.expenseRepository.GetExpenses(command.UserID, command.AccountID, filters, 0, 0)
	if err != nil {
		return err
	}

	if !command.Force && (len(categoryIDs) > 1 || len(expenses) > 0) {
		return errors.New("Category is not empty")
	}

	for _, exp := range expenses {
		h.expenseRepository.DeleteExpense(exp.ID)
	}

	for _, cat := range categoryIDs {
		h.categoryRepository.DeleteCategory(cat)
	}

	return nil
}
