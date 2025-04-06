package request

type CreateCategoryRequest struct {
	Name  string  `json:"name" binding:"required"`
	Color *string `json:"color,omitempty" binding:"omitempty,len=7"`
}

type UpdateCategoryRequest struct {
	Name  *string `json:"name,omitempty"`
	Color *string `json:"color,omitempty" binding:"omitempty,len=7"`
}
