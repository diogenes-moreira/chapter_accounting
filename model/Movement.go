package model

import (
	"gorm.io/gorm"
	"time"
)

type Movement struct {
	gorm.Model
	Amount  float64   `json:"amount"`
	Type    string    `json:"type"`
	Receipt string    `json:"receipt"`
	Date    time.Time `json:"date"`
}

func (m Movement) Credit() bool {
	return IsCredit(m.Type)
}
