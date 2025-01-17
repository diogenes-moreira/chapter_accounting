package model

import (
	"encoding/json"
	"gorm.io/gorm"
)

type RollingBalance struct {
	gorm.Model
	Movements []*Movement   `json:"movements" gorm:"many2many:RollingBalance_Movements;"`
	Adder     MovementAdder `json:"-" gorm:"-"`
}

func (b *RollingBalance) TotalOutcomes() float64 {
	out := 0.0
	for _, movement := range b.Movements {
		if movement.Debit() {
			out += movement.Amount
		}
	}
	return out
}

func (b *RollingBalance) AddMovement(movement *Movement) error {
	b.Movements = append(b.Movements, movement)
	return nil
}

func (b *RollingBalance) Balance(adder MovementAdder) float64 {
	balance := 0.0
	for _, movement := range b.Movements {
		balance = adder.AddMovementTo(balance, movement)
	}
	return balance
}

func (b *RollingBalance) Incomes() []*Movement {
	out := make([]*Movement, 0)
	for _, movement := range b.Movements {
		if movement.Credit() {
			out = append(out, movement)
		}
	}
	return out
}

func (b *RollingBalance) TotalIncomes() float64 {
	total := 0.0
	for _, movement := range b.Incomes() {
		total += movement.Amount
	}
	return total
}

func (b *RollingBalance) TotalExpenses() float64 {
	total := 0.0
	for _, movement := range b.Expenses() {
		total += movement.Amount
	}
	return total
}

func (b *RollingBalance) Expenses() []*Movement {
	out := make([]*Movement, 0)
	for _, movement := range b.Movements {
		if movement.Debit() {
			out = append(out, movement)
		}
	}
	return out
}

func (b *RollingBalance) MarshalJSON() ([]byte, error) {
	type Alias RollingBalance
	return json.Marshal(&struct {
		*Alias
		Balance  float64 `json:"balance"`
		Incomes  float64 `json:"incomes"`
		Outcomes float64 `json:"outcomes"`
	}{
		Alias:    (*Alias)(b),
		Balance:  b.Balance(b.Adder),
		Incomes:  b.TotalIncomes(),
		Outcomes: b.TotalOutcomes(),
	})
}

type MovementAdder interface {
	AddMovement(movement *Movement) error
	AddMovementTo(current float64, movement *Movement) float64
}
