package expense

import "time"

type Expense struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"userId"`
	AccountID   int64     `json:"-"`
	CategoryID  int64     `json:"categoryId"`
	Description string    `json:"description"`
	Amount      int64     `json:"amount"`
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type ExpenseFilters struct {
	MinAmount   *int64
	MaxAmount   *int64
	MinDate     *time.Time
	MaxDate     *time.Time
	CategoryIDs []int64
}

func NewExpense(userID, accountID, categoryID int64, description string, amount int64, date time.Time) *Expense {
	return &Expense{
		UserID:      userID,
		AccountID:   accountID,
		CategoryID:  categoryID,
		Description: description,
		Amount:      amount,
		Date:        date,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
}

func (e *Expense) Update(categoryID int64, description string, amount int64, date time.Time) {
	e.CategoryID = categoryID
	e.Description = description
	e.Amount = amount
	e.Date = date
	e.UpdatedAt = time.Now().UTC()
}
