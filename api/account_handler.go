package api

import "time"

type CreateAccountRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	CurrencyCode string `json:"currencyCode"`
}

type UpdateAccountRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	CurrencyCode string `json:"currencyCode"`
}

type UpdateUserSettingsRequest struct {
	DefaultAccountID int64  `json:"defaultAccountId"`
	Theme            string `json:"theme"`
}

type CreateCategoryRequest struct {
	ParentID    int64  `json:"parentId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type UpdateCategoryRequest struct {
	ParentID    int64  `json:"parentId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type CreateExpenseRequest struct {
	CategoryID  int64     `json:"categoryId"`
	Description string    `json:"description"`
	Amount      int64     `json:"amount"`
	Date        time.Time `json:"date"`
}

type UpdateExpenseRequest struct {
	CategoryID  int64     `json:"categoryId"`
	Description string    `json:"description"`
	Amount      int64     `json:"amount"`
	Date        time.Time `json:"date"`
}

type ListExpensesRequest struct {
	MinAmount *int64     `json:"minAmount"`
	MaxAmount *int64     `json:"maxAmount"`
	MinDate   *time.Time `json:"minDate"`
	MaxDate   *time.Time `json:"maxDate"`
}
