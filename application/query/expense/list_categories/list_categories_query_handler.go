package list_categories

import (
	"github.com/angelbachev/go-money-api/domain/category"
)

type ListCategoriesQueryHandler struct {
	categoryRepository category.CategoryRepository
}

func NewListCategoriesQueryHandler(
	categoryRepository category.CategoryRepository,
) *ListCategoriesQueryHandler {
	return &ListCategoriesQueryHandler{
		categoryRepository: categoryRepository,
	}
}

func (h ListCategoriesQueryHandler) Handle(query ListCategoriesQuery) (*ListCategoriesResponse, error) {
	categories, err := h.categoryRepository.GetCategoryTree(query.AccountID)
	if err != nil {
		return nil, err
	}

	return NewListCategoriesResponse(categories), nil
}
