package update_category

import (
	"github.com/angelbachev/go-money-api/domain/category"
)

type UpdateCategoryCommandHandler struct {
	categoryRepository category.CategoryRepositoryInterface
}

func NewUpdateCategoryCommandHandler(
	categoryRepository category.CategoryRepositoryInterface,
) *UpdateCategoryCommandHandler {
	return &UpdateCategoryCommandHandler{
		categoryRepository: categoryRepository,
	}
}

func (h UpdateCategoryCommandHandler) Handle(command UpdateCategoryCommand) (*UpdateCategoryResponse, error) {
	category, err := h.categoryRepository.GetCategoryByID(command.ID)
	if err != nil {
		return nil, err
	}

	if category == nil || category.AccountID != command.AccountID {
		return nil, err
	}

	if command.ParentID == 0 {
		command.ParentID, err = h.categoryRepository.GetRootCategoryID(command.AccountID)
		if err != nil {
			return nil, err
		}
	}

	category.Update(command.ParentID, command.Name, command.Description, command.Icon)
	err = h.categoryRepository.UpdateCategory(category)
	if err != nil {
		return nil, err
	}

	return NewUpdateCategoryResponse(category), nil
}
