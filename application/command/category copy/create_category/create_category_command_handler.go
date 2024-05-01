package create_category

import (
	"github.com/angelbachev/go-money-api/domain/category"
)

type CreateCategoryCommandHandler struct {
	categoryRepository category.CategoryRepository
}

func NewCreateCategoryCommandHandler(
	categoryRepository category.CategoryRepository,
) *CreateCategoryCommandHandler {
	return &CreateCategoryCommandHandler{
		categoryRepository: categoryRepository,
	}
}

func (h CreateCategoryCommandHandler) Handle(command CreateCategoryCommand) (*CreateCategoryResponse, error) {
	if command.ParentID == 0 {
		parentID, err := h.categoryRepository.GetRootCategoryID(command.AccountID)
		if err != nil {
			return nil, err
		}
		command.ParentID = parentID
	}

	category := category.NewCategory(command.UserID, command.AccountID, command.ParentID, command.Name, command.Description, command.Icon)

	h.categoryRepository.CreateCategory(category)

	return NewCreateCategoryResponse(category), nil
}
