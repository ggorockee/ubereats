package response

import (
	"ubereats/app/core/helper/common"
)

type CreateCategoryOut struct {
	common.CoreResponse
}

type GetAllCategoryOut struct {
	common.CoreResponse
}

type GetCategoryOut struct {
	Ok      bool            `json:"ok"`
	Message string          `json:"message,omitempty"`
	Data    *CategoryResult `json:"data,omitempty"`
	common.PaginationOutput
}

type CategoryResult struct {
	ID          uint               `json:"id"`
	Name        string             `json:"name"`
	CoverImg    string             `json:"cover_img"`
	Restaurants []SimpleRestaurant `json:"restaurants"`
	TotalPages  int                `json:"total_pages"`
}
