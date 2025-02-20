package dto

// CreateCategory는 카테고리 생성 입력 DTO
type CreateCategory struct {
	Name     string `json:"name"`
	CoverImg string `json:"cover_img"`
}

// UpdateCategory는 카테고리 수정 입력 DTO
type UpdateCategory struct {
	Name     string `json:"name,omitempty"`
	CoverImg string `json:"cover_img,omitempty"`
}

// CategoryInput는 카테고리 조회 입력 DTO
type CategoryInput struct {
	Page int `json:"page"`
}
