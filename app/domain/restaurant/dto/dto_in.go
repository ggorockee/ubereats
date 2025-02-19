package dto

type CreateRestaurant struct {
	Name     string `json:"name"`
	CoverImg string `json:"cover_img"`
	Address  string `json:"address"`
}

type UpdateRestaurant struct {
	Name       string `gorm:"column:name" json:"name,omitempty"`
	IsVegan    bool   `gorm:"column:is_vegan" json:"is_vegan,omitempty"`
	Address    string `gorm:"column:address" json:"address,omitempty"`
	OwnersName string `gorm:"column:owners_name" json:"owners_name,omitempty"`
}
