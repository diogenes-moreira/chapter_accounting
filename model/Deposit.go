package model

import (
	"gorm.io/gorm"
	"time"
)

type Deposit struct {
	gorm.Model
	Amount         float64        `json:"amount"`
	Installments   []*Installment `json:"installments" gorm:"foreignKey:DepositID"`
	GenerationDate time.Time      `json:"generation_date"`
	DepositDate    time.Time      `json:"deposit_date"`
}

func (d *Deposit) AddInstallments(installments []*Installment) {
	for _, installment := range installments {
		d.Amount += installment.GreatChapterAmount
		d.Installments = append(d.Installments, installment)
		installment.Deposit = d
	}
}

func (d *Deposit) CreateMovement() (*Movement, error) {
	mt, err := GetDeposit()
	if err != nil {
		return nil, err
	}
	return &Movement{
		Amount:       d.Amount,
		MovementType: mt,
		Description:  "Depósito de en la Tesorería del Capítulo",
		Date:         d.DepositDate,
	}, nil
}

func (d *Deposit) GreatChapterMovement() (*Movement, error) {
	mt, err := GetGreatChapterDeposit()
	if err != nil {
		return nil, err

	}
	return &Movement{
		Amount:       d.Amount,
		MovementType: mt,
		Description:  "Depósito de en la Tesorería del Gran Capítulo",
		Date:         d.DepositDate,
	}, nil
}

func (d *Deposit) In(out []*Deposit) bool {
	if len(out) == 0 {
		return false
	} else {
		for _, deposit := range out {
			if deposit.ID == d.ID {
				return true
			}
		}
		return false
	}
}

func (d *Deposit) AdvancePaymentAmount() float64 {
	out := 0.0
	for _, installment := range d.Installments {
		if installment.InAdvance() {
			out += installment.GreatChapterAmount
		}
	}
	return out
}
