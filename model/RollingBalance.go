package model

import "gorm.io/gorm"

type RollingBalance struct {
	gorm.Model
	Movements []*Movement `json:"movements" gorm:"many2many:RollingBalance_Movements;"`
}

func (b *RollingBalance) AddMovement(movement *Movement) {
	b.Movements = append(b.Movements, movement)
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

type MovementAdder interface {
	AddMovement(movement *Movement)
	AddMovementTo(current float64, movement *Movement) float64
}
