package entity

type Category struct {
	CoreEntity
	Name string `json:"name" gorm:"column:name"`
}

func (m *Category) UpdateDelProperty() {
	m.IsDel = "Y"
}
