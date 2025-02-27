package entity

import "time"

type CoreEntity struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

type CoreResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message,omitempty"`
}
