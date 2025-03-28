package entity

import "time"

type CoreEntity struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	IsDel     string     `json:"is_del" gorm:"column:is_del;size:1;default:'N'"`
}
