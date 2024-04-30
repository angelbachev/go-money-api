package get_category

import (
	"github.com/angelbachev/go-money-api/domain/category"
)

type GetCategoryResponse struct {
	CategoryTree *category.CategoryTree
}

func NewGetCategoryResponse(categoryTree *category.CategoryTree) *GetCategoryResponse {
	return &GetCategoryResponse{
		CategoryTree: categoryTree,
	}
}
