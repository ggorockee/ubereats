package response

// type RestaurantResponse struct {
// 	ID        int                   `json:"id"`         // CoreEntity에서 상속
// 	CreatedAt string                `json:"created_at"` // 시간 형식을 문자열로 변환
// 	UpdatedAt string                `json:"updated_at"` // 시간 형식을 문자열로 변환
// 	Name      string                `json:"name"`
// 	CoverImg  string                `json:"cover_img"`
// 	Address   string                `json:"address"`
// 	Category  *Category             `json:"category,omitempty"` // nullable 반영
// 	Owner     userRes.OwnerResponse `json:"owner"`              // User를 위한 중첩 DTO
// }

// type ListRestaurantRes struct {
// 	ID        int       `json:"id"`         // CoreEntity에서 상속
// 	CreatedAt string    `json:"created_at"` // 시간 형식을 문자열로 변환
// 	UpdatedAt string    `json:"updated_at"` // 시간 형식을 문자열로 변환
// 	Name      string    `json:"name"`
// 	CoverImg  string    `json:"cover_img"`
// 	Address   string    `json:"address"`
// 	Category  *Category `json:"category,omitempty"` // nullable 반영
// }

// // Category는 RestaurantResponse 내에서 사용되는 카테고리 DTO
// type Category struct {
// 	ID       int    `json:"id"`
// 	Name     string `json:"name"`
// 	CoverImg string `json:"cover_img"`
// }

// func GenerateListRestaurantRes(m *[]entity.Restaurant) []ListRestaurantRes {
// 	resp := make([]ListRestaurantRes, len(*m))
// 	for i, mm := range *m {
// 		var category *Category
// 		if mm.Category != nil {
// 			category = &Category{
// 				ID:       mm.Category.ID, // Category.ID 사용
// 				Name:     mm.Category.Name,
// 				CoverImg: mm.Category.CoverImg, // Category.CoverImg 사용
// 			}
// 		}
// 		resp[i] = ListRestaurantRes{
// 			ID:        mm.ID,
// 			CreatedAt: mm.CreatedAt.Format("2006-01-02 15:04:05"),
// 			UpdatedAt: mm.UpdatedAt.Format("2006-01-02 15:04:05"),
// 			Name:      mm.Name,
// 			CoverImg:  mm.CoverImg,
// 			Address:   mm.Address,
// 			Category:  category, // nil 가능
// 		}

// 	}

// 	return resp
// }

// func GenRestaurantRes(m *entity.Restaurant) RestaurantResponse {
// 	resp := RestaurantResponse{
// 		ID:        m.ID,
// 		CreatedAt: m.CreatedAt.Format("2006-01-02 15:04:05"), // 시간 형식 변환
// 		UpdatedAt: m.UpdatedAt.Format("2006-01-02 15:04:05"), // 시간 형식 변환
// 		Name:      m.Name,
// 		CoverImg:  m.CoverImg,
// 		Address:   m.Address,
// 		Owner: userRes.OwnerResponse{
// 			ID:    m.Owner.ID,
// 			Email: m.Owner.Email,
// 		},
// 	}

// 	// Category가 nil이 아닌 경우에만 매핑

// 	if m.Category != nil {

// 		resp.Category = &Category{
// 			ID:       m.Category.ID,
// 			Name:     m.Category.Name,
// 			CoverImg: m.Category.CoverImg,
// 		}
// 	}

// 	return resp
// }
