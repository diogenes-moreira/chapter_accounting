package model

import "gorm.io/gorm"

type Companion struct {
	gorm.Model
	FirstName            string `json:"first_name"`
	Email                string `json:"email"`
	PhoneNumber          string `json:"phone_number"`
	LastNames            string `json:"last_names"`
	IsGreatChapterMember bool   `json:"is_great_chapter"`
}

func (b Companion) InstallmentAmount(monthlyCharge *ChargeType) float64 {
	if b.IsGreatChapterMember {
		return monthlyCharge.GreatChapterAmount
	}
	return monthlyCharge.Amount
}

func (b Companion) GreatChapterAmount(monthlyCharge *ChargeType) float64 {
	return monthlyCharge.GreatChapterAmount
}
