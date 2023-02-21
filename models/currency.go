package models

type Currency struct {
	BaseModel
	Code string `gorm:"size:5" json:"code"`
	Name string `gorm:"size:64" json:"name"`
}

func (Currency) TableName() string {
	return "currencies"
}
