package delete_category

type DeleteCategoryCommand struct {
	ID        int64
	AccountID int64
	UserID    int64
	Force     bool
}
