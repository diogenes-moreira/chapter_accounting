package model

import "gorm.io/gorm"

type ChargeType struct {
	gorm.Model
	Name               string  `json:"name" gorm:"unique"`
	Amount             float64 `json:"amount"`
	GreatChapterAmount float64 `json:"great_chapter_amount"`
}

// TODO resolverlo desde la base de datos
func GetMonthlyCharge() *ChargeType {
	return &ChargeType{
		Name:               "Cuota mensual",
		Amount:             1000,
		GreatChapterAmount: 100,
	}
}
