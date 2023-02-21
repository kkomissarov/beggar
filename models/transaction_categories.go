package models

import "github.com/shopspring/decimal"

type TransactionCategory struct {
	BaseModel
	UserID     int             `json:"-"`
	User       User            `gorm:"foreignKey:UserID" json:"user"`
	Name       string          `gorm:"size:64" json:"name"`
	Type       TransactionType `gorm:"type:transaction_types"`
	MonthLimit decimal.Decimal `gorm:"type:decimal(16, 2);" json:"month_limit"`
}

func (TransactionCategory) TableName() string {
	return "transaction_categories"
}
