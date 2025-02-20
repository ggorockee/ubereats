package dto

// CreateRestaurant는 레스토랑 생성 입력 DTO
type CreateRestaurant struct {
	Name         string `json:"name"`
	CoverImg     string `json:"cover_img"`
	Address      string `json:"address"`
	CategoryName string `json:"category_name"`
}

// UpdateRestaurant는 레스토랑 수정 입력 DTO
type UpdateRestaurant struct {
	Name         string `json:"name,omitempty"`
	CoverImg     string `json:"cover_img,omitempty"`
	Address      string `json:"address,omitempty"`
	CategoryName string `json:"category_name,omitempty"`
}

// DeleteRestaurant는 레스토랑 삭제 입력 DTO
type DeleteRestaurant struct {
	RestaurantID int `json:"restaurant_id"`
}

// RestaurantInput는 특정 레스토랑 조회 입력 DTO
type RestaurantInput struct {
	RestaurantID int `json:"restaurant_id"`
}

// RestaurantsInput는 레스토랑 목록 조회 입력 DTO
type RestaurantsInput struct {
	Page int `json:"page"`
}

// SearchRestaurant는 레스토랑 검색 입력 DTO
type SearchRestaurant struct {
	Query string `json:"query"`
	Page  int    `json:"page"`
}

// CategoryInput는 카테고리 조회 입력 DTO
type CategoryInput struct {
	Slug string `json:"slug"`
	Page int    `json:"page"`
}
