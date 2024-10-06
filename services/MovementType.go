package services

import "argentina-tresury/model"

func CreateMovementType(movementType *model.MovementType) error {
	if err := model.DB.Create(movementType).Error; err != nil {
		return err
	}
	return nil
}

func GetManualMovementTypes() ([]*model.MovementType, error) {
	return model.GetManualMovementTypes()
}

func GetMovementTypes() ([]*model.MovementType, error) {
	return model.GetMovementTypes()
}

func GetMovementType(u uint) (*model.MovementType, error) {
	var movementType model.MovementType
	if err := model.DB.First(&movementType, u).Error; err != nil {
		return nil, err
	}
	return &movementType, nil
}

func UpdateMovementType(m *model.MovementType) error {
	if err := model.DB.Save(m).Error; err != nil {
		return err
	}
	return nil
}

func DeleteMovementType(u uint) error {
	if err := model.DB.Delete(&model.MovementType{}, u).Error; err != nil {
		return err
	}
	return nil
}
