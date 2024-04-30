package update_category

import (
	"github.com/angelbachev/go-money-api/domain/category"
)

type UpdateCategoryResponse struct {
	Category *category.Category
}

func NewUpdateCategoryResponse(category *category.Category) *UpdateCategoryResponse {
	return &UpdateCategoryResponse{
		Category: category,
	}
}
