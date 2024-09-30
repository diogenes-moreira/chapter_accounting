package model

import (
	"gorm.io/gorm"
	"time"
)

type Installment struct {
	gorm.Model
	Month              int      `json:"month"`
	Year               int      `json:"year"`
	Amount             float64  `json:"amount"`
	DueDate            string   `json:"due_date"`
	GrandChapterAmount float64  `json:"grand_chapter_amount"`
	Paid               bool     `json:"paid"`
	DepositID          *uint    `json:"deposit_id"` // Referencia opcional a un depósito
	Deposit            *Deposit `json:"deposit"`    // Puntero a la estructura Deposit
}

func (i *Installment) Apply(a *Affiliation) {
	if !i.Paid && i.Amount > a.Balance {
		i.Paid = true
		a.AddMovement(
			&Movement{
				Type:   InstalmentCancellation,
				Amount: i.Amount,
				Date:   time.Now()})
	}
}
