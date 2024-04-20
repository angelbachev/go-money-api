package models

import "time"

type Expense struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"userId"`
	BudgetID    int64     `json:"budgetId"`
	CategoryID  int64     `json:"categoryId"`
	Description string    `json:"description"`
	Amount      int64     `json:"amount"`
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewExpense(userID, budgetID, categoryID int64, description string, amount int64, date time.Time) *Expense {
	return &Expense{
		UserID:      userID,
		BudgetID:    budgetID,
		CategoryID:  categoryID,
		Description: description,
		Amount:      amount,
		Date:        date,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
}
