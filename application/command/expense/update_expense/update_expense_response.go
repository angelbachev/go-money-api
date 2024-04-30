package update_expense

import (
	"github.com/angelbachev/go-money-api/domain/expense"
)

type UpdateExpenseResponse struct {
	Expense *expense.Expense
}

func NewUpdateExpenseResponse(expense *expense.Expense) *UpdateExpenseResponse {
	return &UpdateExpenseResponse{
		Expense: expense,
	}
}
