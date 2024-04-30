package list_category_icons

import (
	"os"
)

type ListCategoryIconsQueryHandler struct {
}

func NewListCategoryIconsQueryHandler() *ListCategoryIconsQueryHandler {
	return &ListCategoryIconsQueryHandler{}
}

func (h ListCategoryIconsQueryHandler) Handle(query ListCategoryIconsQuery) (*ListCategoryIconsResponse, error) {
	icons, err := os.ReadDir("./files/images/categories")
	if err != nil {
		return nil, err
	}

	var iconNames []string
	for _, icon := range icons {
		iconNames = append(iconNames, icon.Name())
	}

	return NewListCategoryIconsResponse(iconNames), nil
}
