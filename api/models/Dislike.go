package models

import (
	"gorm.io/gorm"
)

type Disike struct {
	gorm.Model
	UserID uint32 `gorm:"not null" json:"user_id"`
	PostID uint32 `gorm:"not null" json:"post_id"`
}
