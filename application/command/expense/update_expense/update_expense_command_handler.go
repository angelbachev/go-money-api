package update_expense

import (
	"github.com/angelbachev/go-money-api/domain/expense"
)

type UpdateExpenseCommandHandler struct {
	expenseRepository expense.ExpenseRepository
}

func NewUpdateExpenseCommandHandler(
	expenseRepository expense.ExpenseRepository,
) *UpdateExpenseCommandHandler {
	return &UpdateExpenseCommandHandler{
		expenseRepository: expenseRepository,
	}
}

func (h UpdateExpenseCommandHandler) Handle(command UpdateExpenseCommand) (*UpdateExpenseResponse, error) {
	expense, err := h.expenseRepository.GetExpenseByID(command.UserID, command.AccountID, command.ID)
	if err != nil {
		return nil, err
	}

	if expense == nil || expense.AccountID != command.AccountID {
		return nil, err
	}

	expense.Update(command.CategoryID, command.Description, command.Amount, command.Date)
	err = h.expenseRepository.UpdateExpense(expense)
	if err != nil {
		return nil, err
	}

	return NewUpdateExpenseResponse(expense), nil
}
