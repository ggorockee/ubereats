package response

import "ubereats/app/core/entity"

// BaseResponse는 기본 응답 구조
type BaseResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}

// CategoryResponse는 단일 카테고리 응답 DTO
type CategoryResponse struct {
	BaseResponse
	ID          int                 `json:"id"`
	CreatedAt   string              `json:"created_at"`
	UpdatedAt   string              `json:"updated_at"`
	Name        string              `json:"name"`
	CoverImg    string              `json:"cover_img"`
	Restaurants []entity.Restaurant `json:"restaurants,omitempty"`
}

// AllCategoriesResponse는 모든 카테고리 응답 DTO
type AllCategoriesResponse struct {
	BaseResponse
	Categories []CategoryWithCount `json:"categories,omitempty"`
}

// CategoryWithCount는 카테고리와 레스토랑 수를 포함한 DTO
type CategoryWithCount struct {
	ID              int    `json:"id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	Name            string `json:"name"`
	CoverImg        string `json:"cover_img"`
	RestaurantCount int    `json:"restaurant_count"`
}

func GenCategoryRes(m *entity.Category) CategoryResponse {
	return CategoryResponse{
		BaseResponse: BaseResponse{Ok: true},
		ID:           m.ID,
		CreatedAt:    m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    m.UpdatedAt.Format("2006-01-02 15:04:05"),
		Name:         m.Name,
		CoverImg:     m.CoverImg,
		Restaurants:  m.Restaurants,
	}
}

func GenCategoriesRes(categories *[]entity.Category, counts []int) []CategoryWithCount {
	resp := make([]CategoryWithCount, len(*categories))
	for i, cate := range *categories {
		resp[i] = CategoryWithCount{
			ID:              cate.ID,
			CreatedAt:       cate.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:       cate.UpdatedAt.Format("2006-01-02 15:04:05"),
			Name:            cate.Name,
			CoverImg:        cate.CoverImg,
			RestaurantCount: counts[i],
		}
	}
	return resp
}
