package model

import "gorm.io/gorm"

type ChargeType struct {
	gorm.Model
	Code               string   `json:"code" gorm:"unique"`
	Name               string   `json:"name"`
	Amount             float64  `json:"amount"`
	GreatChapterAmount float64  `json:"great_chapter_amount"`
	ChapterID          *uint    `json:"chapter_id"`
	Chapter            *Chapter `json:"-"`
}

const monthlyChargeCode = "monthly_charge"
const exaltationChargeCode = "exaltation_charge"

func GetMonthlyCharge(chapter *Chapter) *ChargeType {
	return GetChargeType(chapter, monthlyChargeCode)
}

func GetChargeType(chapter *Chapter, code string) *ChargeType {
	for _, chargeType := range chapter.ChargeTypes {
		if chargeType.Code == code {
			return chargeType
		}
	}
	return nil

}

func GetExaltationCharge(chapter *Chapter) *ChargeType {
	return GetChargeType(chapter, exaltationChargeCode)
}

func InitChargeTypes(chapter *Chapter) []*ChargeType {
	return []*ChargeType{
		{
			Code:               monthlyChargeCode,
			Name:               "Cuota Mensual",
			Amount:             1000,
			GreatChapterAmount: 100,
			Chapter:            chapter,
		},
		{
			Code:               exaltationChargeCode,
			Name:               "Exaltación",
			Amount:             5000,
			GreatChapterAmount: 500,
			Chapter:            chapter,
		},
	}
}
