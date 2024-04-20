package models

import "time"

type Category struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"userId"`
	BudgetID    int64     `json:"budgetId"`
	ParentID    int64     `json:"categoryId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CategoryTree struct {
	ID          int64           `json:"id"`
	UserID      int64           `json:"userId"`
	ParentID    int64           `json:"-"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
	Children    []*CategoryTree `json:"children"`
}

func NewCategory(userID, budgetID, parentID int64, name, description string) *Category {
	return &Category{
		UserID:      userID,
		BudgetID:    budgetID,
		ParentID:    parentID,
		Name:        name,
		Description: description,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
}
