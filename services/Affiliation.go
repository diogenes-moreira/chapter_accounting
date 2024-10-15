package services

import (
	"argentina-tresury/model"
	"gorm.io/gorm"
	"time"
)

func GetAffiliation(id uint) (*model.Affiliation, error) {
	affiliation := &model.Affiliation{}
	if err := model.DB.Preload("Chapter").
		Preload("Chapter.TreasurerRollingBalance").
		Preload("Chapter.TreasurerRollingBalance.Movements").
		Preload("Chapter.TreasurerRollingBalance.Movements.MovementType").
		Preload("RollingBalance").
		Preload("RollingBalance.Movements").
		Preload("RollingBalance.Movements.MovementType").
		Preload("Installments", func(db *gorm.DB) *gorm.DB {
			return db.Order("Installments.DUE_DATE")
		}).
		Preload("Companion").
		First(affiliation, id).Error; err != nil {
		return nil, err
	}
	return affiliation, nil
}

func CreatePayment(affiliationId uint, amount float64, receipt string, date time.Time) error {
	affiliation, err := GetAffiliation(affiliationId)
	if err != nil {
		return err
	}
	mt, err := model.GetCapitationPayment()
	if err != nil {
		return err
	}
	mov := &model.Movement{
		MovementType: mt,
		Amount:       amount,
		Receipt:      receipt,
		Date:         date,
		Description:  "Pago de Hermano " + affiliation.Companion.FirstName + " " + affiliation.Companion.LastNames,
	}
	err = affiliation.AddMovement(mov)
	if err != nil {
		return err
	}
	return model.DB.Save(affiliation).Error
}

func CreateAffiliationExpense(affiliationId uint, amount float64, receipt string, date time.Time,
	expenseType string, description string) error {
	affiliation, err := GetAffiliation(affiliationId)
	if err != nil {
		return err
	}
	mt, err := model.GetMovementType(expenseType)
	if err != nil {
		return err
	}
	mov := &model.Movement{
		MovementType: mt,
		Amount:       amount,
		Receipt:      receipt,
		Date:         date,
		Description:  description,
	}
	err = affiliation.AddMovement(mov)
	if err != nil {
		return err
	}
	mt, err = model.GetCapitationPayment()
	if err != nil {
		return err
	}
	mov = &model.Movement{
		MovementType: mt,
		Amount:       amount,
		Receipt:      receipt,
		Date:         date,
		Description:  "Pago de Hermano " + affiliation.Companion.FirstName + " " + affiliation.Companion.LastNames,
	}
	err = affiliation.Chapter.AddMovement(mov)

	if err != nil {
		return err
	}

	return model.DB.Save(affiliation).Error
}
