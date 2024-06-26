package get_category

import (
	"github.com/angelbachev/go-money-api/domain/category"
)

type GetCategoryQueryHandler struct {
	categoryRepository category.CategoryRepository
}

func NewGetCategoryQueryHandler(
	categoryRepository category.CategoryRepository,
) *GetCategoryQueryHandler {
	return &GetCategoryQueryHandler{
		categoryRepository: categoryRepository,
	}
}

func (h GetCategoryQueryHandler) Handle(query GetCategoryQuery) (*GetCategoryResponse, error) {
	category, err := h.categoryRepository.GetSingleCategoryTree(query.ID)
	if err != nil {
		return nil, err
	}
	// TODO: check user and account

	return NewGetCategoryResponse(category), nil
}
