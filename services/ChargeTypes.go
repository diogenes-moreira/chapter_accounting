package services

import "argentina-tresury/model"

func GetChargeTypes() ([]*model.ChargeType, error) {
	var out []*model.ChargeType
	if model.DB.Find(&out).Error != nil {
		return nil, model.DB.Error
	}
	return out, nil
}

func UpdateChargeType(ct *model.ChargeType) error {
	if model.DB.Save(ct).Error != nil {
		return model.DB.Error
	}
	return nil
}

func UpdateChargeById(id uint, amount float64, amountGreatChapter float64) error {
	ct, err := GetChargeType(id)
	if err != nil {
		return err
	}
	ct.Amount = amount
	ct.GreatChapterAmount = amountGreatChapter
	return UpdateChargeType(ct)
}

func GetChargeType(id uint) (*model.ChargeType, error) {
	var ct model.ChargeType
	if model.DB.First(&ct, id).Error != nil {
		return nil, model.DB.Error
	}
	return &ct, nil
}
