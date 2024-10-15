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
const affiliationChargeCode = "affiliation_charge"

func GetMonthlyCharge(chapter *Chapter) (*ChargeType, error) {
	return GetChargeType(chapter, monthlyChargeCode)
}

func GetChargeType(chapter *Chapter, code string) (*ChargeType, error) {
	var ct ChargeType
	if err := DB.Where("chapter_id = ? AND code = ?", chapter.ID, code).First(&ct).Error; err != nil {
		return nil, err
	}
	return &ct, nil
}

func GetExaltationCharge(chapter *Chapter) (*ChargeType, error) {
	return GetChargeType(chapter, exaltationChargeCode)
}

func GetAffiliationCharge(chapter *Chapter) (*ChargeType, error) {
	return GetChargeType(chapter, affiliationChargeCode)
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
		{
			Code:               affiliationChargeCode,
			Name:               "Afiliación",
			Amount:             1000,
			GreatChapterAmount: 100,
			Chapter:            chapter,
		},
	}
}
