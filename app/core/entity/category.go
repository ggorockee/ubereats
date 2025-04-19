package entity

type Category struct {
	CoreEntity
	Name        string       `gorm:"column:name" json:"name" validate:"required,min=5" mapstructure:"name"`
	CoverImg    string       `gorm:"column:cover_img" json:"coverImg" validate:"required" mapstructure:"cover_img"`
	Restaurants []Restaurant `gorm:"foreignKey:CategoryRefer" json:"restaurants,omitempty" mapstructure:"restaurants"`
}

func (m *Category) UpdateDelProperty() {
	m.IsDel = "Y"
}
