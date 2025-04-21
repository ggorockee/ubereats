package dto

import "ubereats/app/core/helper/common"

type CreateCategoryIn struct {
	Name     string `json:"name" validate:"required" mapstructure:"name"`
	CoverImg string `json:"cover_img" validate:"required" mapstructure:"cover_img"`
}

type GetCategoryIn struct {
	common.PaginationParams
	Name string `json:"name" validate:"required" mapstructure:"name"`
}
