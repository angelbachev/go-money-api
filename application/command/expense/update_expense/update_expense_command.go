package update_expense

import "time"

type UpdateExpenseCommand struct {
	ID          int64
	UserID      int64
	AccountID   int64
	CategoryID  int64     `json:"categoryId"`
	Description string    `json:"description"`
	Amount      int64     `json:"amount"`
	Date        time.Time `json:"date"`
}
