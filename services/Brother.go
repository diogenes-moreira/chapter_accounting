package services

import (
	"argentina-tresury/model"
)

func CreateBrother(brother *model.Brother) error {
	if err := model.DB.Create(brother).Error; err != nil {
		return err
	}
	return nil
}

func GetBrothers() ([]model.Brother, error) {
	var brothers []model.Brother
	if err := model.DB.Find(&brothers).Error; err != nil {
		return nil, err
	}
	return brothers, nil
}

func GetBrother(u uint) (*model.Brother, error) {
	var brother model.Brother
	if err := model.DB.First(&brother, u).Error; err != nil {
		return nil, err
	}
	return &brother, nil
}

func UpdateBrother(m *model.Brother) error {
	if err := model.DB.Save(m).Error; err != nil {
		return err
	}
	return nil
}

func CreateExaltation(brother *model.Brother, isHonorary bool, chapter *model.Chapter) error {
	if err := model.DB.Create(&brother).Error; err != nil {
		return err
	}
	affiliation, err := CreateAffiliation(brother, chapter, isHonorary)
	if err != nil {
		return err
	}
	if !isHonorary {
		charge := model.GetExaltationCharge(chapter)
		affiliation.AddCharge(charge)
		if err := model.DB.Save(affiliation).Error; err != nil {
			return err
		}
	}
	return nil
}
