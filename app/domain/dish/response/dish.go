package response

import "ubereats/app/core/entity"

type BaseResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}

// CreateDishResponse는 요리 생성 응답 DTO
type CreateDishResponse struct {
	BaseResponse
}

// UpdateDishResponse는 요리 수정 응답 DTO
type UpdateDishResponse struct {
	BaseResponse
}

// DeleteDishResponse는 요리 삭제 응답 DTO
type DeleteDishResponse struct {
	BaseResponse
}

// DishResponse는 특정 요리 응답 DTO
type DishResponse struct {
	BaseResponse
	Dish *entity.Dish `json:"dish,omitempty"`
}

// DishsResponse는 요리 목록 응답 DTO
type DishsResponse struct {
	BaseResponse
	Results      []entity.Dish `json:"results,omitempty"`
	TotalPages   int           `json:"total_pages"`
	TotalResults int           `json:"total_results"`
}

// SearchDishResponse는 요리 검색 응답 DTO
type SearchDishResponse struct {
	BaseResponse
	Dishs        []entity.Dish `json:"dishs,omitempty"`
	TotalPages   int           `json:"total_pages"`
	TotalResults int           `json:"total_results"`
}

// AllCategoriesResponse는 모든 카테고리 응답 DTO (Dish와 공유)
type AllCategoriesResponse struct {
	BaseResponse
	Categories []entity.Category `json:"categories,omitempty"`
}
