package request_schema

import "github.com/shopspring/decimal"

type CreateAccountBody struct {
	Name       string          `json:"name" binding:"required"`
	Balance    decimal.Decimal `json:"balance"`
	CurrencyID int             `json:"currency_id" binding:"required"`
}
