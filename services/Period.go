package services

import (
	"argentina-tresury/model"
)

func CreatePeriod(period *model.Period) error {
	if err := model.DB.Create(period).Error; err != nil {
		return err
	}
	return nil
}

func GetPeriods() ([]model.Period, error) {
	var periods []model.Period
	if err := model.DB.Find(&periods).Error; err != nil {
		return nil, err
	}
	return periods, nil
}

func GetPeriod(u uint) (*model.Period, error) {
	var period model.Period
	if err := model.DB.First(&period, u).Error; err != nil {
		return nil, err
	}
	return &period, nil
}

func UpdatePeriod(m *model.Period) error {
	if err := model.DB.Save(m).Error; err != nil {
		return err
	}
	return nil
}
