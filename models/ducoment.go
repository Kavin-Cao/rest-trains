package models

type Document struct{
	RTModel
	Markdown string	`gorm:"Type:longtext;not null"`
	Title string `gorm:"Size:50;not null"`
	Creator uint
}