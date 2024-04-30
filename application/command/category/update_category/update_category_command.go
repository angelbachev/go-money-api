package update_category

type UpdateCategoryCommand struct {
	ID          int64
	AccountID   int64
	UserID      int64
	ParentID    int64  `json:"parentId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
