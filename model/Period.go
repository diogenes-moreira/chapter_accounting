package model

import "gorm.io/gorm"

type Period struct {
	gorm.Model
	Year int `json:"year"`
}
