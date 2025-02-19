package dto

type CreateRestaurant struct {
	Name       string `json:"name,omitempty"`        //@required
	CoverImg   string `json:"cover_img,omitempty"`   //@required
	Address    string `json:"address,omitempty"`     //@required
	CategoryID int    `json:"category_id,omitempty"` //@required
	OwnerID    int    `json:"owner_id,omitempty"`    //@required
}

type UpdateRestaurant struct {
	Name       string `json:"name,omitempty"`        //@required
	CoverImg   string `json:"cover_img,omitempty"`   //@required
	Address    string `json:"address,omitempty"`     //@required
	CategoryID int    `json:"category_id,omitempty"` //@required
}
