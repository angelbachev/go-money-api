package list_expenses

import "time"

type ListExpensesQuery struct {
	UserID      int64
	AccountID   int64
	MinAmount   *int64
	MaxAmount   *int64
	MinDate     *time.Time
	MaxDate     *time.Time
	CategoryIDs []int64
	Page        int64
	Limit       int64
}
