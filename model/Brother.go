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

func (b Brother) InstallmentAmount(chapter *Chapter) float64 {
	getMonthlyCharge := GetMonthlyCharge(chapter)

	if b.IsGrandChapterMember {
		return getMonthlyCharge.GreatChapterAmount
	}
	return getMonthlyCharge.Amount
}

func (b Brother) GreatChapterAmount(chapter *Chapter) float64 {
	return GetMonthlyCharge(chapter).GreatChapterAmount
}
