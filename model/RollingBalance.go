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

type MovementAdder interface {
	AddMovement(movement *Movement)
	AddMovementTo(current float64, movement *Movement) float64
}
