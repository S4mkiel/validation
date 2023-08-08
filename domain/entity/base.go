package entity

import "time"

type Base struct {
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt time.Time `json:"deleted_at" gorm:"default:CURRENT_TIMESTAMP"`
}
