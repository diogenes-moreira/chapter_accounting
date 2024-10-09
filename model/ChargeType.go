package model

import "gorm.io/gorm"

type ChargeType struct {
	gorm.Model
	Code               string  `json:"code" gorm:"unique"`
	Name               string  `json:"name"`
	Amount             float64 `json:"amount"`
	GreatChapterAmount float64 `json:"great_chapter_amount"`
}

const monthlyChargeCode = "monthly_charge"
const exaltationChargeCode = "exaltation_charge"

func GetMonthlyCharge() (*ChargeType, error) {
	var charge ChargeType
	err := DB.Where("code = ?", monthlyChargeCode).First(&charge).Error
	if err != nil {
		return nil, err
	}
	return &charge, nil
}

func GetExaltationCharge() (*ChargeType, error) {
	var charge ChargeType
	err := DB.Where("code = ?", exaltationChargeCode).First(&charge).Error
	if err != nil {
		return nil, err
	}
	return &charge, nil
}

func InitChargeTypes() {
	monthlyCharge := ChargeType{
		Code:               monthlyChargeCode,
		Name:               "Monthly Charge",
		Amount:             100.0,
		GreatChapterAmount: 150.0,
	}
	exaltationCharge := ChargeType{
		Code:               exaltationChargeCode,
		Name:               "Exaltation Charge",
		Amount:             150.0,
		GreatChapterAmount: 200.0,
	}
	DB.Create(&monthlyCharge)
	DB.Create(&exaltationCharge)
}
