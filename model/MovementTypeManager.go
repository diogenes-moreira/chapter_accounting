package model

import (
	"errors"
	"gorm.io/gorm"
)

func GetManualMovementTypes() ([]*MovementType, error) {
	var out []*MovementType
	movementTypes, err := GetMovementTypes()
	if err != nil {
		return nil, err
	}
	for _, movementType := range movementTypes {
		if movementType.Manual {
			out = append(out, movementType)
		}
	}
	return out, nil
}

func GetMovementTypes() ([]*MovementType, error) {
	var movementTypes []*MovementType
	if err := DB.Find(&movementTypes).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			InitMovementTypes()
			return GetMovementTypes()
		}
		return nil, err
	}
	if len(movementTypes) == 0 {
		InitMovementTypes()
		return GetMovementTypes()
	}
	return movementTypes, nil
}

func GetMovementType(name string) (*MovementType, error) {
	var movementType MovementType
	if err := DB.First(&movementType, "Code = ?", name).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			InitMovementTypes()
			return GetMovementType(name)
		}
		return nil, err
	}
	return &movementType, nil
}

func InitMovementTypes() {
	DB.Create(&MovementType{Code: _capitationPayment, Description: "Pago de cuota de capita", Expense: false,
		Credit: true, Manual: false})
	DB.Create(&MovementType{Code: _instalmentCancellation, Description: "Cancelación de cuota", Expense: false,
		Credit: false, Manual: false})
	DB.Create(&MovementType{Code: _grandChapterMonthlyDebit, Description: "Débito mensual de Gran Capítulo",
		Expense: false, Credit: false, Manual: false})
	DB.Create(&MovementType{Code: _adjustmentOfInstallments, Description: "Ajuste de cuotas", Expense: false,
		Credit: false, Manual: false})
	DB.Create(&MovementType{Code: _bagIncome, Description: "Ingreso de bolsa", Expense: false, Credit: true,
		Manual: true})
	DB.Create(&MovementType{Code: _brotherIncome, Description: "Ingreso de hermano", Expense: false, Credit: true,
		Manual: true})
	DB.Create(&MovementType{Code: _deposit, Description: "Depósito", Expense: false, Credit: false, Manual: true})
	DB.Create(&MovementType{Code: _greatChapterDeposit, Description: "Depósito de Gran Capítulo", Expense: false,
		Credit: true, Manual: false})
	DB.Create(&MovementType{Code: _initialAmount, Description: "Monto inicial", Expense: false, Credit: true,
		Manual: false})
	DB.Create(&MovementType{Code: _stationery_items, Description: "Artículos de librería", Expense: true, Credit: false,
		Manual: true})
	DB.Create(&MovementType{Code: _agape, Description: "Ágape", Expense: true, Credit: false, Manual: true})
	DB.Create(&MovementType{Code: _donation, Description: "Donación", Expense: false, Credit: true, Manual: true})
	DB.Create(&MovementType{Code: _other, Description: "Otro gasto", Expense: true, Credit: false, Manual: true})
	DB.Create(&MovementType{Code: _otherIncome, Description: "Otro ingreso", Expense: false, Credit: true,
		Manual: true})
}

func GetCapitationPayment() (*MovementType, error) {
	return GetMovementType(_capitationPayment)
}

func GetMovementTypeBrotherIncome() (*MovementType, error) {
	return GetMovementType(_brotherIncome)
}

func GetInstalmentCancellation() (*MovementType, error) {
	return GetMovementType(_instalmentCancellation)
}

func GetBagIncome() (*MovementType, error) {
	return GetMovementType(_bagIncome)
}

func GetAdjustmentOfInstallments() (*MovementType, error) {
	return GetMovementType(_adjustmentOfInstallments)
}

func GetGrandChapterMonthlyDebit() (*MovementType, error) {
	return GetMovementType(_grandChapterMonthlyDebit)
}

func GetDeposit() (*MovementType, error) {
	return GetMovementType(_deposit)
}

func GetGreatChapterDeposit() (*MovementType, error) {
	return GetMovementType(_greatChapterDeposit)
}

const _capitationPayment = "capitation_payment"
const _instalmentCancellation = "instalment_cancellation"
const _grandChapterMonthlyDebit = "grand_chapter_monthly_debit"
const _adjustmentOfInstallments = "adjustment_of_installments"
const _bagIncome = "bag_income"
const _brotherIncome = "brother_income"
const _deposit = "deposit"
const _greatChapterDeposit = "great_chapter_deposit"
const _initialAmount = "initial_amount"
const _stationery_items = "stationery_items"
const _agape = "agape"
const _donation = "donation"
const _other = "other expense"
const _otherIncome = "other income"
