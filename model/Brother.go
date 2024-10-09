package model

import "gorm.io/gorm"

type Brother struct {
	gorm.Model
	FirstName            string `json:"first_name"`
	Email                string `json:"email"`
	PhoneNumber          string `json:"phone_number"`
	LastNames            string `json:"last_names"`
	IsGrandChapterMember bool   `json:"is_grand_chapter"`
}

func (b Brother) InstallmentAmount() float64 {
	getMonthlyCharge, _ := GetMonthlyCharge()
	if b.IsGrandChapterMember {
		return getMonthlyCharge.GreatChapterAmount
	}
	return getMonthlyCharge.Amount
}

func (b Brother) GreatChapterAmount() float64 {
	getMonthlyCharge, _ := GetMonthlyCharge()
	return getMonthlyCharge.GreatChapterAmount
}
