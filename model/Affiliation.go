package model

import (
	"encoding/json"
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
	Installments     []*Installment `json:"installments" gorm:"foreignKey:AffiliationID"`
	StartDate        time.Time      `json:"start_date"`
	EndDate          *time.Time     `json:"end_date"` // Puede ser nulo si la afiliación está activa
	RollingBalance   RollingBalance `json:"rolling_balance"`
	RollingBalanceID uint           `json:"rolling_balance_id"`
	Balance          float64        `json:"balance" gorm:"-"`
	Honorary         bool           `json:"honorary"`
}

func (a *Affiliation) AddMovement(movement *Movement) {
	a.Chapter.AddMovement(movement)
	a.RollingBalance.AddMovement(movement)
	a.Balance = a.RollingBalance.Balance(a)
	a.ApplyInstallments()
}

func (a *Affiliation) AddMovementTo(current float64, movement *Movement) float64 {
	if movement.Credit() || movement.Expense() {
		return current + movement.Amount
	}
	return current - movement.Amount
}

func (a *Affiliation) ApplyInstallments() error {
	for _, installment := range a.Installments {
		err := installment.Apply()
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Affiliation) AddInstallment(installment *Installment) {
	installment.Affiliation = a
	a.Installments = append(a.Installments, installment)
}

func (a *Affiliation) NetBalance() float64 {
	a.Balance = a.RollingBalance.Balance(a)
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

func (a *Affiliation) BrotherName() string {
	return a.Brother.FirstName + " " + a.Brother.LastNames
}

func (a *Affiliation) AddCharge(charge *ChargeType) {
	if charge != nil && !a.Honorary {
		a.Installments = append(a.Installments, &Installment{
			OnTheSpot:          true,
			Description:        charge.Name,
			Amount:             charge.Amount,
			DueDate:            time.Now(),
			GrandChapterAmount: charge.GreatChapterAmount,
			Paid:               false,
		})
	}
}

func (a *Affiliation) SetChapter(c *Chapter) {
	a.Chapter = c
	if !a.Honorary {
		for _, installment := range c.PeriodPendingInstallments(a.Brother) {
			a.AddInstallment(installment)
		}
	}
}

func (a *Affiliation) IsPeriod(p uint) bool {
	return *a.PeriodID == p
}

func (a *Affiliation) MarshalJSON() ([]byte, error) {
	type Alias Affiliation
	return json.Marshal(&struct {
		*Alias
		Overdue float64 `json:"overdue"`
	}{
		Alias:   (*Alias)(a),
		Overdue: a.NetBalance(),
	})
}
