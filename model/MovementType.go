package model

import "gorm.io/gorm"

type MovementType struct {
	gorm.Model
	Code        string `json:"name" gorm:"unique"`
	Description string `json:"description"`
	Expense     bool   `json:"expense"`
	Credit      bool   `json:"credit"`
	Manual      bool   `json:"manual"`
}
