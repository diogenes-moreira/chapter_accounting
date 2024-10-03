package model

import "gorm.io/gorm"

type ChargeType struct {
	gorm.Model
	Name               string  `json:"name" gorm:"unique"`
	Amount             float64 `json:"amount"`
	GreatChapterAmount float64 `json:"great_chapter_amount"`
}

func AffiliationCharge() *Movement {
	return nil // TODO
}
