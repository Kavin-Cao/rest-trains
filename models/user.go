package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Mobile string `gorm:"Size:16;not null"`
	Passwd string `gorm:"Size:50;not null"`
	Avatar string
	Nickname string `gorm:"Size:20"`
}

type Project struct {
	gorm.Model
	Creator uint
	Name string `gorm:"Size:16;not null"`
	Logo string
}