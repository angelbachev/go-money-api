package delete_expense

import (
	"github.com/angelbachev/go-money-api/domain/expense"
)

type DeleteExpenseCommandHandler struct {
	expenseRepository expense.ExpenseRepositoryInterface
}

func NewDeleteExpenseCommandHandler(
	expenseRepository expense.ExpenseRepositoryInterface,
) *DeleteExpenseCommandHandler {
	return &DeleteExpenseCommandHandler{
		expenseRepository: expenseRepository,
	}
}

func (h DeleteExpenseCommandHandler) Handle(command DeleteExpenseCommand) error {
	// TODO: validate that the user owns the account and category

	_, err := h.expenseRepository.GetExpenseByID(command.UserID, command.AccountID, command.ID)
	if err != nil {
		return err
	}

	return h.expenseRepository.DeleteExpense(command.ID)
}
