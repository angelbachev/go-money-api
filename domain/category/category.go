package category

import "time"

type Category struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"userId"`
	AccountID   int64     `json:"-"`
	ParentID    int64     `json:"categoryId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CategoryTree struct {
	ID          int64           `json:"id"`
	UserID      int64           `json:"userId"`
	ParentID    int64           `json:"-"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Icon        string          `json:"icon"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
	Children    []*CategoryTree `json:"children"`
}

func NewCategory(userID, accountID, parentID int64, name, description, icon string) *Category {
	return &Category{
		UserID:      userID,
		AccountID:   accountID,
		ParentID:    parentID,
		Name:        name,
		Description: description,
		Icon:        icon,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
}

func (c *Category) Update(parentID int64, name, description, icon string) {
	c.ParentID = parentID
	c.Name = name
	c.Description = description
	c.Icon = icon
	c.UpdatedAt = time.Now().UTC()
}
