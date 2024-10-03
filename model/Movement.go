package model

import (
	"gorm.io/gorm"
	"time"
)

type Movement struct {
	gorm.Model
	Amount      float64   `json:"amount"`
	Type        string    `json:"type"`
	Receipt     string    `json:"receipt"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

func (m Movement) Credit() bool {
	return IsCredit(m.Type)
}

func (m Movement) Debit() bool {
	return !m.Credit()
}
