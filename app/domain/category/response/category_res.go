package response

import (
	"ubereats/app/core/entity"
)

type CategoryResponse struct {
	entity.CoreEntity
	Name     string `json:"name"`
	CoverImg string `json:"cover_img"`
}

func GenCategoriesRes(m *[]entity.Category) []CategoryResponse {
	resp := make([]CategoryResponse, len(*m))
	for i, cate := range *m {
		resp[i] = GenCategoryRes(&cate)
	}
	return resp
}

func GenCategoryRes(m *entity.Category) CategoryResponse {
	resp := CategoryResponse{
		CoreEntity: entity.CoreEntity{
			ID:        m.ID,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		},
		Name:     m.Name,
		CoverImg: m.CoverImg,
	}

	return resp
}
