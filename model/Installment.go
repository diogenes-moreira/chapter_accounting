package model

import (
	"gorm.io/gorm"
	"time"
)

type Installment struct {
	gorm.Model
	Month              int          `json:"month"`
	Year               int          `json:"year"`
	OnTheSpot          bool         `json:"on_the_spot"`
	Description        string       `json:"description"`
	Amount             float64      `json:"amount"`
	DueDate            time.Time    `json:"due_date"`
	GrandChapterAmount float64      `json:"grand_chapter_amount"`
	Paid               bool         `json:"paid"`
	DepositID          *uint        `json:"deposit_id"` // Referencia opcional a un depósito
	Deposit            *Deposit     `json:"deposit"`    // Puntero a la estructura Deposit
	AffiliationID      uint         `json:"affiliation_id"`
	Affiliation        *Affiliation `json:"affiliation"`
}

func (i *Installment) Apply() error {
	if !i.Paid && i.Amount > i.Affiliation.Balance {
		mt, err := GetInstalmentCancellation()
		if err != nil {
			return err
		}
		i.Paid = true
		i.Affiliation.AddMovement(
			&Movement{
				MovementType: mt,
				Amount:       i.Amount,
				Description:  "Cancelación de " + i.Description,
				Date:         time.Now()})
	}
	return nil
}

func (i *Installment) IsDue() bool {
	return time.Now().After(i.DueDate)
}

func (i *Installment) Pending() bool {
	return i.Paid && i.Deposit == nil
}

func (i *Installment) DueAt(period *Period, month uint) bool {
	return i.Month == int(month) && i.Year == period.Year
}
