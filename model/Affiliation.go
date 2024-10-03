package model

import (
	"gorm.io/gorm"
	"time"
)

type Affiliation struct {
	gorm.Model
	PeriodID         *uint          `json:"period_id"`
	Period           *Period        `json:"period"`
	BrotherID        *uint          `json:"brother_id"`
	Brother          *Brother       `json:"brother"`
	ChapterID        *uint          `json:"chapter_id"`
	Chapter          *Chapter       `json:"chapter"`
	Installments     []*Installment `json:"installments" gorm:"foreignKey:AfiliacionID"`
	StartDate        time.Time      `json:"start_date"`
	EndDate          *time.Time     `json:"end_date"` // Puede ser nulo si la afiliación está activa
	RollingBalance   RollingBalance `json:"rolling_balance"`
	RollingBalanceID uint           `json:"rolling_balance_id"`
	Balance          float64        `json:"balance"`
	Honorary         bool           `json:"honorary"`
}

func (a *Affiliation) AddMovement(movement *Movement) {
	a.Chapter.AddMovement(movement)
	a.RollingBalance.AddMovement(movement)
	a.Balance = a.RollingBalance.Balance(a)
	a.ApplyInstallments()
}

func (a *Affiliation) AddMovementTo(current float64, movement *Movement) float64 {
	if movement.Credit() {
		return current + movement.Amount
	}
	return current - movement.Amount
}

func (a *Affiliation) ApplyInstallments() {
	for _, installment := range a.Installments {
		installment.Apply()
	}
}

func (a *Affiliation) AddInstallment(installment *Installment) {
	installment.Affiliation = a
	a.Installments = append(a.Installments, installment)
}

func (a *Affiliation) NetBalance() float64 {
	out := 0.0
	for _, installment := range a.Installments {
		if !installment.Paid && installment.IsDue() {
			out += installment.Amount
		}
	}
	return out + a.Balance
}

func (a *Affiliation) PendingInstallments() []*Installment {
	out := make([]*Installment, 0)
	for _, installment := range a.Installments {
		if installment.Pending() {
			out = append(out, installment)
		}
	}
	return out

}

func (a *Affiliation) GrandChapterAmountDueAt(month uint) float64 {
	out := 0.0
	for _, installment := range a.Installments {
		if installment.DueAt(a.Period, month) {
			out += installment.GrandChapterAmount
		}
	}
	return out
}

func (a *Affiliation) UpdateInstallment(amount float64, greatChapterAmount float64) float64 {
	out := 0.0
	for _, installment := range a.Installments {
		if !installment.Paid {
			installment.Amount = amount
			out += greatChapterAmount - installment.GrandChapterAmount
			installment.GrandChapterAmount = greatChapterAmount
		}
	}
	return out
}
