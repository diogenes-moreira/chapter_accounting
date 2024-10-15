package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName  string   `json:"user_name" gorm:"unique"`
	Password  string   `json:"password"`
	Chapter   *Chapter `json:"-"`
	ChapterID *uint    `json:"chapter_id"`
	Profile   string   `json:"profile"`
}

func (u User) IsAdmin() bool {
	return u.Profile == "admin"
}

func (u User) IsTreasurer() bool {
	return u.Profile == "treasurer"
}
