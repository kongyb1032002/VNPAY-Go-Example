package model

import "time"

type Entity struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime"`
	// DeletedAt time.Time `json:"deleted_at" gorm:"autoCreateTime"`
	Status uint8 `json:"status" gorm:"default:1"`
}
