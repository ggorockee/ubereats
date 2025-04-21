package common

type PaginationParams struct {
	Page  int `json:"page" query:"page" default:"1" validate:"min=1"`
	Limit int `json:"limit" query:"limit"`
}

type PaginationOutput struct {
	TotalPages *int `json:"total_pages,omitempty"`
}
