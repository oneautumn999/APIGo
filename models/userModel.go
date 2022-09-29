package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	Username string
}

type Photo struct {
	gorm.Model
	Title    string
	Caption  string
	PhotoUrl string
	UserID   uint
	User     User
}
