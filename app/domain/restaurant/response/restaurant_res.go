package response

import (
	"ubereats/app/core/entity"
	userRes "ubereats/app/domain/user/response"
)

type RestaurantResponse struct {
	ID        int                   `json:"id"`         // CoreEntity에서 상속
	CreatedAt string                `json:"created_at"` // 시간 형식을 문자열로 변환
	UpdatedAt string                `json:"updated_at"` // 시간 형식을 문자열로 변환
	Name      string                `json:"name"`
	CoverImg  string                `json:"cover_img"`
	Address   string                `json:"address"`
	Category  *Category             `json:"category,omitempty"` // nullable 반영
	Owner     userRes.OwnerResponse `json:"owner"`              // User를 위한 중첩 DTO
}

// Category는 RestaurantResponse 내에서 사용되는 카테고리 DTO입니다.
type Category struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	CoverImg string `json:"cover_img"`
}

// OwnerResponse는 RestaurantResponse 내에서 사용되는 소유자 DTO입니다.

func GenRestaurantRes(m *entity.Restaurant) RestaurantResponse {
	resp := RestaurantResponse{
		ID:        m.ID,
		CreatedAt: m.CreatedAt.Format("2006-01-02 15:04:05"), // 시간 형식 변환
		UpdatedAt: m.UpdatedAt.Format("2006-01-02 15:04:05"), // 시간 형식 변환
		Name:      m.Name,
		CoverImg:  m.CoverImg,
		Address:   m.Address,
		Owner: userRes.OwnerResponse{
			ID:    m.Owner.ID,
			Email: m.Owner.Email,
		},
	}

	// Category가 nil이 아닌 경우에만 매핑

	if m.Category != nil {

		resp.Category = &Category{
			ID:       m.Category.ID,
			Name:     m.Category.Name,
			CoverImg: m.Category.CoverImg,
		}
	}

	return resp
}
