package models

import "time"

type RevokedToken struct {
	BaseModel
	Token     string    `gorm:"unique;index;"`
	ExpiredAt time.Time `gorm:"notnull;index;"`
}

func (RevokedToken) TableName() string {
	return "revoked_tokens"
}
