package api

import "time"

type CreateBudgetRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateCategoryRequest struct {
	ParentID    int64  `json:"parentId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateExpenseRequest struct {
	CategoryID  int64     `json:"categoryId"`
	Description string    `json:"description"`
	Amount      int64     `json:"amount"`
	Date        time.Time `json:"date"`
}
