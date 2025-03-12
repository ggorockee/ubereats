package common

type PaginationInput struct {
	Page int `json:"page" default:"1"`
}

type PaginationOutput struct {
	BaseResponse
	TotalPages   *int `json:"total_pages,omitempty"`
	TotalResults *int `json:"total_results,omitempty"`
}
