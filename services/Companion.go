package services

import (
	"argentina-tresury/model"
)

func CreateCompanion(companion *model.Companion) error {
	if err := model.DB.Create(companion).Error; err != nil {
		return err
	}
	return nil
}

func GetCompanions() ([]model.Companion, error) {
	var companions []model.Companion
	if err := model.DB.Find(&companions).Error; err != nil {
		return nil, err
	}
	return companions, nil
}

func GetCompanion(u uint) (*model.Companion, error) {
	var companion model.Companion
	if err := model.DB.First(&companion, u).Error; err != nil {
		return nil, err
	}
	return &companion, nil
}

func UpdateCompanion(m *model.Companion) error {
	if err := model.DB.Save(m).Error; err != nil {
		return err
	}
	return nil
}

func CreateExaltation(companion *model.Companion, isHonorary bool, chapter *model.Chapter) error {
	if err := model.DB.Create(&companion).Error; err != nil {
		return err
	}
	affiliation, err := CreateAffiliation(companion, chapter, isHonorary)
	if err != nil {
		return err
	}
	if !isHonorary {

		charge, err := model.GetExaltationCharge(chapter)
		if err != nil {
			return err
		}
		affiliation.AddCharge(charge)
		if err = model.DB.Save(affiliation).Error; err != nil {
			return err
		}
	}
	return nil
}

func CreateCompanionAffiliation(companion *model.Companion, isHonorary bool, chapter *model.Chapter) error {
	affiliation, err := CreateAffiliation(companion, chapter, isHonorary)
	if err != nil {
		return err
	}
	if !isHonorary {
		charge, err := model.GetAffiliationCharge(chapter)
		if err != nil {
			return err
		}
		affiliation.AddCharge(charge)
		if err := model.DB.Save(affiliation).Error; err != nil {
			return err
		}
	}
	return nil
}
