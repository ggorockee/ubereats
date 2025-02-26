package response

import (
	"ubereats/app/core/entity"
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

func ResponseCategoriesOutput(categories *[]entity.Category) []restDto.Category {
	categoryOutput := make([]restDto.Category, len(*categories))
	for i, category := range *categories {
		c := restDto.Category{
			ID:       category.ID,
			Name:     category.Name,
			CoverImg: category.CoverImg,
		}

		categoryOutput[i] = c
	}

	return categoryOutput
}
