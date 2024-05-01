package list_expenses

type ListExpensesResponse struct {
	Items      []*ExpenseResponse `json:"items"`
	TotalCount int64              `json:"totalCount"`
}

func NewListExpensesResponse(expenses []*ExpenseResponse, totalCount int64) *ListExpensesResponse {
	return &ListExpensesResponse{
		Items:      expenses,
		TotalCount: totalCount,
	}
}
