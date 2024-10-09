package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName  string   `json:"user_name" gorm:"unique"`
	Password  string   `json:"password"`
	IsAdmin   bool     `json:"is_admin"`
	Chapter   *Chapter `json:"chapter"`
	ChapterID *uint    `json:"chapter_id"`
}
