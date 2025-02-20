package response

import (
	"ubereats/app/core/entity"

	restDto "ubereats/app/domain/restaurant/dto"
)

// CreateRestaurantResponse는 레스토랑 생성 응답 DTO
type CreateRestaurantOutput struct {
	BaseResponse
}

// UpdateRestaurantResponse는 레스토랑 수정 응답 DTO
type EditRestaurantOutput struct {
	BaseResponse
}

// DeleteRestaurantResponse는 레스토랑 삭제 응답 DTO
type DeleteRestaurantOutput struct {
	BaseResponse
}

type RestaurantsOutput struct {
	BaseResponse
	Results      []restDto.Restaurant `json:"results,omitempty"`
	TotalPages   int                  `json:"totalPages"`
	TotalResults int                  `json:"totalResults"`
}

type RestaurantOutput struct {
	BaseResponse
	Restaurant *restDto.Restaurant `json:"restaurant,omitempty"`
}

type SearchRestaurantOutput struct {
	BaseResponse
	Restaurants  []restDto.Restaurant `json:"restaurants,omitempty"`
	TotalResults int                  `json:"totalResults"`
	TotalPages   int                  `json:"totalPages"`
}

// AllCategoriesResponse는 모든 카테고리 응답 DTO
type AllCategoriesResponse struct {
	BaseResponse
	Categories []entity.Category `json:"categories,omitempty"`
}

// CategoryResponse는 카테고리 조회 응답 DTO
type CategoryResponse struct {
	BaseResponse
	Restaurants []entity.Restaurant `json:"restaurants,omitempty"`
	Category    *entity.Category    `json:"category,omitempty"`
	TotalPages  int                 `json:"total_pages"`
}

// RestaurantsResponse는 레스토랑 목록 응답 DTO
type RestaurantsResponse struct {
	BaseResponse
	Results      []entity.Restaurant `json:"results,omitempty"`
	TotalPages   int                 `json:"total_pages"`
	TotalResults int                 `json:"total_results"`
}

// RestaurantResponse는 특정 레스토랑 응답 DTO
type RestaurantResponse struct {
	BaseResponse
	Restaurant *entity.Restaurant `json:"restaurant,omitempty"`
}

// SearchRestaurantResponse는 레스토랑 검색 응답 DTO
type SearchRestaurantResponse struct {
	BaseResponse
	Restaurants  []entity.Restaurant `json:"restaurants,omitempty"`
	TotalPages   int                 `json:"total_pages"`
	TotalResults int                 `json:"total_results"`
}

// CategoryWithCount는 카테고리와 레스토랑 수를 포함한 DTO
type CategoryWithCount struct {
	BaseResponse
	ID              uint   `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	Name            string `json:"name"`
	CoverImg        string `json:"cover_img"`
	RestaurantCount int64  `json:"restaurant_count"`
}
