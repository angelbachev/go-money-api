package list_category_icons

type ListCategoryIconsResponse struct {
	Icons []string
}

func NewListCategoryIconsResponse(icons []string) *ListCategoryIconsResponse {
	return &ListCategoryIconsResponse{
		Icons: icons,
	}
}
