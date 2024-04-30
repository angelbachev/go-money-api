package create_expense

import "time"

type CreateExpenseCommand struct {
	UserID      int64
	AccountID   int64
	CategoryID  int64     `json:"categoryId"`
	Description string    `json:"description"`
	Amount      int64     `json:"amount"`
	Date        time.Time `json:"date"`
}
