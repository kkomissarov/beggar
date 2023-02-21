package models

import (
	"github.com/shopspring/decimal"
)

type Account struct {
	BaseModel
	Name       string          `gorm:"size:64;not null" json:"name"`
	Balance    decimal.Decimal `gorm:"type:decimal(16, 2);default:0;not null;" json:"balance"`
	UserID     int             `gorm:"not null" json:"-"`
	User       User            `gorm:"foreignKey:UserID" json:"user"`
	CurrencyID int             `gorm:"not null" json:"-"`
	Currency   Currency        `gorm:"foreignKey:CurrencyID" json:"currency"`
}

func (Account) TableName() string {
	return "accounts"
}
