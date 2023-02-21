package models

type User struct {
	BaseModel
	Email    string `gorm:"size:64;unique" json:"email"`
	Password string `gorm:"size:128" json:"-"`
}

func (User) TableName() string {
	return "users"
}
