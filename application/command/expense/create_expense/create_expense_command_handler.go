package create_expense

import (
	"github.com/angelbachev/go-money-api/domain/expense"
)

type CreateExpenseCommandHandler struct {
	expenseRepository expense.ExpenseRepositoryInterface
}

func NewCreateExpenseCommandHandler(
	expenseRepository expense.ExpenseRepositoryInterface,
) *CreateExpenseCommandHandler {
	return &CreateExpenseCommandHandler{
		expenseRepository: expenseRepository,
	}
}

func (h CreateExpenseCommandHandler) Handle(command CreateExpenseCommand) (*CreateExpenseResponse, error) {
	// TODO: validate that the user owns the account and category and category exists

	expense := expense.NewExpense(command.UserID, command.AccountID, command.CategoryID, command.Description, command.Amount, command.Date)
	if err := h.expenseRepository.CreateExpense(expense); err != nil {
		return nil, err
	}

	return NewCreateExpenseResponse(expense), nil
}
