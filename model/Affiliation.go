package model

import (
	"gorm.io/gorm"
	"time"
)

type Affiliation struct {
	gorm.Model
	BrotherID        uint           `json:"brother_id"`
	Brother          Brother        `json:"brother"`
	ChapterID        uint           `json:"chapter_id"`
	Chapter          Chapter        `json:"chapter"`
	Installments     []Installment  `json:"installments" gorm:"foreignKey:AfiliacionID"`
	StartDate        time.Time      `json:"start_date"`
	EndDate          *time.Time     `json:"end_date"` // Puede ser nulo si la afiliación está activa
	RollingBalance   RollingBalance `json:"rolling_balance"`
	RollingBalanceID uint           `json:"rolling_balance_id"`
	Balance          float64        `json:"balance"`
}

func (a *Affiliation) AddMovement(movement *Movement) {
	a.RollingBalance.AddMovement(movement)
	a.Balance = a.RollingBalance.Balance(a)
	a.ApplyInstallments()
	a.Chapter.AddMovement(movement)
}

func (a *Affiliation) AddMovementTo(current float64, movement *Movement) float64 {
	if movement.Credit() {
		return current + movement.Amount
	}
	return current - movement.Amount
}

func (a *Affiliation) ApplyInstallments() {
	for _, installment := range a.Installments {
		installment.Apply(a)
	}
}
