package response

import (
	"ubereats/app/core/entity"
	"ubereats/app/core/helper/common"
)

type CreateRestaurantOut struct {
	common.CoreResponse
}

type EditRestaurantOut struct {
	common.CoreResponse
}

type DeleteRestaurantOut struct {
	common.CoreResponse
}

type AllRestaurantOut struct {
	Ok           bool               `json:"ok"`
	Message      string             `json:"message,omitempty"`
	Results      []SimpleRestaurant `json:"results,omitempty"`
	TotalPages   int                `json:"totalPages"`
	TotalResults int                `json:"totalResults"`
}

type SimpleRestaurant struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	CoverImg   string `json:"cover_img"`
	Address    string `json:"address"`
	CategoryID uint   `json:"category_id"`
	OwnerID    uint   `json:"owner_id"`
}

func ToSimpleRestaurants(restaurants []entity.Restaurant) []SimpleRestaurant {
	result := make([]SimpleRestaurant, len(restaurants))
	for i, r := range restaurants {
		result[i] = SimpleRestaurant{
			ID:         r.ID,
			Name:       r.Name,
			CoverImg:   r.CoverImg,
			Address:    r.Address,
			CategoryID: *r.CategoryRefer, // 또는 r.CategoryID
			OwnerID:    r.OwnerRefer,     // 또는 r.OwnerID
		}
	}
	return result
}
