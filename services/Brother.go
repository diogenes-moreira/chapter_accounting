package services

import (
	"argentina-tresury/db"
	"argentina-tresury/model"
)

func CreateBrother(brother *model.Brother) error {
	if err := db.DB.Create(brother).Error; err != nil {
		return err
	}
	return nil
}

func GetBrothers() ([]model.Brother, error) {
	var brothers []model.Brother
	if err := db.DB.Find(&brothers).Error; err != nil {
		return nil, err
	}
	return brothers, nil
}

func GetBrother(u uint) (*model.Brother, error) {
	var brother model.Brother
	if err := db.DB.First(&brother, u).Error; err != nil {
		return nil, err
	}
	return &brother, nil
}

func UpdateBrother(m *model.Brother) error {
	if err := db.DB.Save(m).Error; err != nil {
		return err
	}
	return nil
}
