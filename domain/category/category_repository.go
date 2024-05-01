package category

type CategoryRepository interface {
	CreateCategory(category *Category) error
	GetCategories(accountID int64) ([]int64, error)
	GetCategoryNames(accountID int64) (map[string]int64, error)
	GetCategoryByID(id int64) (*Category, error)
	GetCategoryTree(accountID int64) (*CategoryTree, error)
	GetSingleCategoryTree(id int64) (*CategoryTree, error)
	GetListCategoryIDsAndTheirSubcategories(ids []int64) ([]int64, error)
	DeleteCategory(id int64) error
	UpdateCategory(category *Category) error
	GetRootCategoryID(accountID int64) (int64, error)
}
