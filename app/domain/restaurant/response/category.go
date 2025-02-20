package response

import (
	restDto "ubereats/app/domain/restaurant/dto"
)

type CategoryOutput struct {
	BaseResponse
	Restaurants []restDto.Restaurant `json:"restaurants,omitempty"`
	Category    *restDto.Category    `json:"category,omitempty"`
	TotalPages  int                  `json:"totalPages"`
}

type AllCategoriesOutput struct {
	BaseResponse
	Categories []restDto.Category `json:"categories,omitempty"`
}
