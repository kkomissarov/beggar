package models

import "database/sql/driver"

type TransactionType string

const (
	Income  TransactionType = "income"
	Expense TransactionType = "expense"
)

func (t *TransactionType) Scan(value interface{}) error {
	*t = TransactionType(value.([]byte))
	return nil
}

func (t TransactionType) Value() (driver.Value, error) {
	return string(t), nil
}
