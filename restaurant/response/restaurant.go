package response

import (
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/response"
)

type Restaurant struct {
	Name     string `json:"name"`
	CoverImg string `json:"cover_img"`
	Address  string `json:"address"`
	//CategoryID int    `json:"category_id"` // 외래 키 (nullable)
	//Category   *Category `json:"category,omitempty"` // @ManyToOne, nullable, onDelete: SET NULL
	//OwnerID    int       `json:"owner_id"`                        // 외래 키 (User와 연결)
	//Owner      User      `json:"owner"`                 // @ManyToOne
	//Menu       []Dish    `json:"menu,omitempty"`
}

// CreateRestaurantResponse는 레스토랑 생성 응답 DTO
type CreateRestaurantOutput struct {
	response.BaseResponse
	*Restaurant `json:"restaurant,omitempty"`
}

// UpdateRestaurantResponse는 레스토랑 수정 응답 DTO
type EditRestaurantOutput struct {
	response.BaseResponse
	*Restaurant `json:"restaurant,omitempty"`
}

// DeleteRestaurantResponse는 레스토랑 삭제 응답 DTO
type DeleteRestaurantOutput struct {
	response.BaseResponse
}

type RestaurantsOutput struct {
	response.BaseResponse
	Results      []Restaurant `json:"results,omitempty"`
	TotalPages   int          `json:"totalPages"`
	TotalResults int          `json:"totalResults"`
}

type RestaurantOutput struct {
	response.BaseResponse
	Restaurant *Restaurant `json:"restaurant,omitempty"`
}

type SearchRestaurantOutput struct {
	response.BaseResponse
	Restaurants  []Restaurant `json:"restaurants,omitempty"`
	TotalResults int          `json:"totalResults"`
	TotalPages   int          `json:"totalPages"`
}

// AllCategoriesResponse는 모든 카테고리 응답 DTO
type AllCategoriesResponse struct {
	response.BaseResponse
	Categories []entity.Category `json:"categories,omitempty"`
}

// CategoryResponse는 카테고리 조회 응답 DTO
type CategoryResponse struct {
	response.BaseResponse
	Restaurants []entity.Restaurant `json:"restaurants,omitempty"`
	Category    *entity.Category    `json:"category,omitempty"`
	TotalPages  int                 `json:"total_pages"`
}

// RestaurantsResponse는 레스토랑 목록 응답 DTO
type RestaurantsResponse struct {
	response.BaseResponse
	Results      []entity.Restaurant `json:"results,omitempty"`
	TotalPages   int                 `json:"total_pages"`
	TotalResults int                 `json:"total_results"`
}

// RestaurantResponse는 특정 레스토랑 응답 DTO
type RestaurantResponse struct {
	response.BaseResponse
	Restaurant *entity.Restaurant `json:"restaurant,omitempty"`
}

// SearchRestaurantResponse는 레스토랑 검색 응답 DTO
type SearchRestaurantResponse struct {
	response.BaseResponse
	Restaurants  []entity.Restaurant `json:"restaurants,omitempty"`
	TotalPages   int                 `json:"total_pages"`
	TotalResults int                 `json:"total_results"`
}

// CategoryWithCount는 카테고리와 레스토랑 수를 포함한 DTO
type CategoryWithCount struct {
	response.BaseResponse
	ID              uint   `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	Name            string `json:"name"`
	CoverImg        string `json:"cover_img"`
	RestaurantCount int64  `json:"restaurant_count"`
}
