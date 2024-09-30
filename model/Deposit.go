package model

import (
	"gorm.io/gorm"
	"time"
)

type Deposit struct {
	gorm.Model
	Amount         float64       `json:"amount"`
	Installments   []Installment `json:"installments" gorm:"foreignKey:DepositID"`
	GenerationDate time.Time     `json:"generation_date"`
	DepositDate    time.Time     `json:"deposit_date"`
}
