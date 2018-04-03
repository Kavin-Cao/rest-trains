package models

import "github.com/jinzhu/gorm"

type Document struct{
	gorm.Model
	Markdown string	`gorm:"Type:longtext;not null"`
	Title string `gorm:"Size:50;not null"`
	Creator uint
}