package models

import (
	"github.com/hashicorp/go-uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	Name     string
	Password string
	Salt     string `gorm:"char(6)"`
	Email    string
	IsAdmin  int8
	Status   int8
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Salt, _ = uuid.GenerateUUID()

	return nil
}
