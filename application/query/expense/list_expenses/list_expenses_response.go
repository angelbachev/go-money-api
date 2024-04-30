package list_expenses

import (
	"github.com/angelbachev/go-money-api/domain/expense"
)

type ListExpensesResponse struct {
	Items      []*expense.Expense `json:"items"`
	TotalCount int64              `json:"totalCount"`
}

func NewListExpensesResponse(expenses []*expense.Expense, totalCount int64) *ListExpensesResponse {
	return &ListExpensesResponse{
		Items:      expenses,
		TotalCount: totalCount,
	}
}
