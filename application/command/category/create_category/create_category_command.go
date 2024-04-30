package create_category

type CreateCategoryCommand struct {
	UserID      int64
	AccountID   int64
	ParentID    int64  `json:"parentId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
