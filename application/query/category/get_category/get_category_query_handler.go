package get_category

import (
	"github.com/angelbachev/go-money-api/domain/category"
)

type GetCategoryQueryHandler struct {
	categoryRepository category.CategoryRepositoryInterface
}

func NewGetCategoryQueryHandler(
	categoryRepository category.CategoryRepositoryInterface,
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
