package dto

type UpdateRestaurant struct {
	Name         string `json:"name,omitempty"`
	CoverImg     string `json:"cover_img,omitempty"`
	Address      string `json:"address,omitempty"`
	CategoryName string `json:"category_name,omitempty"`
}
