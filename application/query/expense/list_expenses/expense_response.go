package list_expenses

import (
	"github.com/angelbachev/go-money-api/domain/expense"
)

type ExpenseResponse struct {
	*expense.Expense
	CategoryName string `json:"categoryName"`
}

func NewExpenseResponse(expense *expense.Expense, categoryName string) *ExpenseResponse {
	return &ExpenseResponse{
		Expense:      expense,
		CategoryName: categoryName,
	}
}
