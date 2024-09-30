package model

import "gorm.io/gorm"

type Brother struct {
	gorm.Model
	Name string `json:"name"`
}
