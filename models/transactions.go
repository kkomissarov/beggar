package models

import "github.com/shopspring/decimal"

type Transaction struct {
	BaseModel
	Amount     decimal.Decimal     `json:"amount"`
	Type       TransactionType     `json:"transaction_type"`
	CategoryID int                 `json:"-"`
	Category   TransactionCategory `gorm:"foreignKey:CategoryID" json:"category"`
	AccountID  int                 `json:"-"`
	Account    Account             `gorm:"foreignKey:AccountID" json:"account"`
	Comment    string              `json:"comment"`
}

func (Transaction) TableName() string {
	return "transactions"
}
