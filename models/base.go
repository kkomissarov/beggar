package models

import (
	"time"
)

type BaseModel struct {
	ID        uint       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time  `gorm:"autoCreateTime;type:timestamp;default:NOW()" json:"-"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime;type:timestamp;default:NOW()" json:"-"`
	DeletedAt *time.Time `gorm:"type:timestamp;index" json:"-"`
}
