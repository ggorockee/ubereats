package entity

type Restaurant struct {
	CoreEntity
	Name     string `gorm:"column:name" json:"name" validate:"required,min=5" mapstructure:"name"`
	CoverImg string `gorm:"column:cover_img" json:"coverImg" validate:"required" mapstructure:"cover_img"`
	Address  string `gorm:"column:address;default:'강남'" json:"address" validate:"required" mapstructure:"address"`

	CategoryRefer *uint    `gorm:"column:category_id" json:"category_id" mapstructure:"category_id"` // FK 필드
	Category      Category `gorm:"foreignKey:CategoryRefer;constraint:OnDelete:SET NULL" json:"category,omitempty" mapstructure:"category"`

	OwnerRefer uint `gorm:"column:owner_id" json:"owner_id" mapstructure:"owner_id"` // FK 필드
	Owner      User `gorm:"foreignKey:OwnerRefer;constraint:OnDelete:CASCADE" json:"owner,omitempty" mapstructure:"owner"`
}
