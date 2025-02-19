package response

import "ubereats/app/core/entity"

type BaseResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}

type CreateRestaurantResponse struct {
	BaseResponse
}

type UpdateRestaurantResponse struct {
	BaseResponse
}

type DeleteRestaurantResponse struct {
	BaseResponse
}

type AllCategoriesResponse struct {
	BaseResponse
	Categories []entity.Category `json:"categories,omitempty"`
}

type CategoryResponse struct {
	BaseResponse
	Restaurants []entity.Restaurant `json:"restaurants,omitempty"`
	Category    *entity.Category    `json:"category,omitempty"`
	TotalPages  int                 `json:"total_pages"`
}

type RestaurantsResponse struct {
	BaseResponse
	Results      []entity.Restaurant `json:"results,omitempty"`
	TotalPages   int                 `json:"total_pages"`
	TotalResults int                 `json:"total_results"`
}

type RestaurantResponse struct {
	BaseResponse
	Restaurant *entity.Restaurant `json:"restaurant,omitempty"`
}

type SearchRestaurantResponse struct {
	BaseResponse
	Restaurants  []entity.Restaurant `json:"restaurants,omitempty"`
	TotalPages   int                 `json:"total_pages"`
	TotalResults int                 `json:"total_results"`
}

type CategoryWithCount struct {
	ID              uint   `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	Name            string `json:"name"`
	CoverImg        string `json:"cover_img"`
	RestaurantCount int64  `json:"restaurant_count"`
}
