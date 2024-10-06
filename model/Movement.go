package model

import (
	"gorm.io/gorm"
	"time"
)

type Movement struct {
	gorm.Model
	MovementTypeID *uint         `json:"movement_type_id"`
	MovementType   *MovementType `json:"movement_type"`
	Amount         float64       `json:"amount"`
	Receipt        string        `json:"receipt"`
	Date           time.Time     `json:"date"`
	Description    string        `json:"description"`
}

func (m *Movement) Credit() bool {
	return m.MovementType.Credit
}

func (m *Movement) Debit() bool {
	return !m.Credit()
}

func (m *Movement) Expense() bool {
	return m.MovementType.Expense
}
