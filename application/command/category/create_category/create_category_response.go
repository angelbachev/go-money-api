package create_category

import (
	"github.com/angelbachev/go-money-api/domain/category"
)

type CreateCategoryResponse struct {
	Category *category.Category
}

func NewCreateCategoryResponse(category *category.Category) *CreateCategoryResponse {
	return &CreateCategoryResponse{
		Category: category,
	}
}
