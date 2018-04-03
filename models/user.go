package models

type User struct {
	RTModel
	Mobile string `gorm:"Size:16;not null"`
	Passwd string `gorm:"Size:50;not null"`
	Avatar string
	Nickname string `gorm:"Size:20"`
}

type Project struct {
	Creator uint
	Name string `gorm:"Size:16;not null"`
	Logo string
}