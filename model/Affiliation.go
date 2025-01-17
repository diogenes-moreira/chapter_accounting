package model

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Affiliation struct {
	gorm.Model
	PeriodID         *uint          `json:"period_id"`
	Period           *Period        `json:"-"`
	CompanionID      *uint          `json:"companion_id"`
	Companion        *Companion     `json:"companion"`
	ChapterID        *uint          `json:"chapter_id"`
	Chapter          *Chapter       `json:"-"`
	Installments     []*Installment `json:"installments" gorm:"foreignKey:AffiliationID"`
	StartDate        time.Time      `json:"start_date"`
	EndDate          *time.Time     `json:"end_date"` // Puede ser nulo si la afiliación está activa
	RollingBalance   RollingBalance `json:"-"`
	RollingBalanceID uint           `json:"rolling_balance_id"`
	Balance          float64        `json:"balance" gorm:"-"`
	Honorary         bool           `json:"honorary"`
}

func (a *Affiliation) AddMovement(movement *Movement) error {
	err := a.Chapter.AddMovement(movement)
	if err != nil {
		return err
	}
	err = a.RollingBalance.AddMovement(movement)
	if err != nil {
		return err
	}
	a.Balance = a.RollingBalance.Balance(a)
	return a.ApplyInstallments()
}

func (a *Affiliation) AddMovementTo(current float64, movement *Movement) float64 {
	if movement.Credit() || movement.Expense() {
		return current + movement.Amount
	}
	return current - movement.Amount
}

func (a *Affiliation) ApplyInstallments() error {
	for _, installment := range a.Installments {
		// this is for gorm not maintain a single instance of the relationship bilateral
		installment.Affiliation = a
		out, err := installment.Apply()
		if err != nil {
			return err
		}
		if out {
			break
		}
	}
	return nil
}

func (a *Affiliation) AddInstallment(installment *Installment) {
	installment.Affiliation = a
	a.Installments = append(a.Installments, installment)
}

func (a *Affiliation) OverDue() float64 {
	out := 0.0
	for _, installment := range a.DueInstallments() {
		out += installment.Amount
	}
	return out
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

func (a *Affiliation) GreatChapterAmountDueAt(month uint) float64 {
	out := 0.0
	for _, installment := range a.Installments {
		if installment.DueAt(a.Period, month) {
			out += installment.GreatChapterAmount
		}
	}
	return out
}

func (a *Affiliation) UpdateInstallment(amount float64, greatChapterAmount float64) {
	for _, installment := range a.Installments {
		if !installment.Paid {
			installment.Amount = amount
			installment.GreatChapterAmount = greatChapterAmount
		}
	}
}

func (a *Affiliation) CompanionName() string {
	return a.Companion.FirstName + " " + a.Companion.LastNames
}

func (a *Affiliation) AddCharge(charge *ChargeType) {
	if charge != nil && !a.Honorary {
		a.Installments = append(a.Installments, &Installment{
			OnTheSpot:          true,
			Description:        charge.Name,
			Amount:             charge.Amount,
			DueDate:            time.Now(),
			GreatChapterAmount: charge.GreatChapterAmount,
			Paid:               false,
			ChargeType:         charge,
		})
	}
}

func (a *Affiliation) SetChapter(c *Chapter) error {
	a.Chapter = c
	installments, err := c.PeriodPendingInstallments(a.Companion)
	if err != nil {
		return err
	}
	if !a.Honorary {
		for _, installment := range installments {
			a.AddInstallment(installment)
		}
	}
	return nil
}

func (a *Affiliation) IsPeriod(p uint) bool {
	return *a.PeriodID == p
}

func (a *Affiliation) MarshalJSON() ([]byte, error) {
	type Alias Affiliation
	a.RollingBalance.Adder = a
	a.Balance = a.RollingBalance.Balance(a)
	return json.Marshal(&struct {
		*Alias
		Overdue   float64     `json:"overdue"`
		Movements []*Movement `json:"movements"`
	}{
		Alias:     (*Alias)(a),
		Overdue:   a.OverDue(),
		Movements: a.RollingBalance.Movements,
	})
}

func (a *Affiliation) IsCurrent() bool {
	return a.Period.Current
}

func (a *Affiliation) DueInstallments() []*Installment {
	out := make([]*Installment, 0)
	for _, installment := range a.Installments {
		if installment.IsDue() {
			out = append(out, installment)
		}
	}
	return out
}

func (a *Affiliation) Deposits() []*Deposit {
	out := make([]*Deposit, 0)
	for _, installment := range a.Installments {
		if installment.Deposit != nil {
			out = append(out, installment.Deposit)
		}
	}
	return out
}

func (a *Affiliation) OverDueGreatChapter() float64 {
	out := 0.0
	for _, installment := range a.DueInstallments() {
		if installment.IsDue() {
			out += installment.GreatChapterAmount
		}
	}
	return out
}

func (a *Affiliation) UpdateInstallments(amount float64, amount2 float64, id uint) error {
	for _, installment := range a.Installments {
		if !installment.Paid && installment.ChargeTypeID == id {
			installment.Amount = amount
			installment.GreatChapterAmount = amount2
		}
	}
	return nil

}
