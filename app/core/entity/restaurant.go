package entity

type Restaurant struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Name         string `gorm:"column:name" json:"name"`
	IsVegan      bool   `gorm:"column:is_vegan" json:"is_vegan"`
	Address      string `gorm:"column:address" json:"address"`
	OwnersName   string `gorm:"column:owners_name" json:"owners_name"`
	CategoryName string `gorm:"column:category_name" json:"category_name"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}
