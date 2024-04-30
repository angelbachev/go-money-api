package create_expense

import (
	"github.com/angelbachev/go-money-api/domain/expense"
)

type CreateExpenseResponse struct {
	Expense *expense.Expense
}

func NewCreateExpenseResponse(expense *expense.Expense) *CreateExpenseResponse {
	return &CreateExpenseResponse{
		Expense: expense,
	}
}
