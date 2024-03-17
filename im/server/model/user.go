package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UUID     string `gorm:"column:uuid;not null;unique;size:255"`
	Name     string `gorm:"not null;unique;size:255"`
	Password string `gorm:"not null;size:255"`
}

func (User) TableName() string {
	return "im_user"
}
