package list_categories

import (
	"github.com/angelbachev/go-money-api/domain/category"
)

type ListCategoriesResponse struct {
	CategoryTree *category.CategoryTree
}

func NewListCategoriesResponse(categoryTree *category.CategoryTree) *ListCategoriesResponse {
	return &ListCategoriesResponse{
		CategoryTree: categoryTree,
	}
}
