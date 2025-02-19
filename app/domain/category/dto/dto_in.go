package dto

type CreateCategory struct {
	Name     string `json:"name"`
	CoverImg string `json:"cover_img"`
}

type UpdateCategory struct {
	Name     string `json:"name,omitempty"`
	CoverImg string `json:"cover_img,omitempty"`
}
