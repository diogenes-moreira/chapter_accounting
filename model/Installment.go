package model

import (
	"encoding/json"
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
	GreatChapterAmount float64      `json:"great_chapter_amount"`
	PaidDate           time.Time    `json:"paid_date"`
	Paid               bool         `json:"paid"`
	DepositID          *uint        `json:"deposit_id"`
	Deposit            *Deposit     `json:"deposit"`
	AffiliationID      uint         `json:"affiliation_id"`
	Affiliation        *Affiliation `json:"-"`
	ChargeTypeID       uint         `json:"charge_type_id"`
	ChargeType         *ChargeType  `json:"-"`
}

func (i *Installment) Apply() (bool, error) {
	if i.Paid {
		return false, nil
	}
	if i.Amount <= i.Affiliation.Balance {
		mt, err := GetInstalmentCancellation()
		if err != nil {
			return false, err
		}
		i.Paid = true
		i.PaidDate = time.Now()
		err = i.Affiliation.AddMovement(
			&Movement{
				MovementType: mt,
				Amount:       i.Amount,
				Description:  "Cancelación de " + i.Description,
				Date:         time.Now()})
		if err != nil {
			return false, err
		}
		return true, nil
	} else {
		return true, nil
	}
}

func (i *Installment) IsDue() bool {
	return time.Now().After(i.DueDate) && !i.Paid
}

func (i *Installment) Pending() bool {
	return i.Paid && i.Deposit == nil
}

func (i *Installment) DueAt(period *Period, month uint) bool {
	return i.Month == int(month) && i.Year == period.Year
}

func (i *Installment) MarshalJSON() ([]byte, error) {
	type Alias Installment
	return json.Marshal(&struct {
		*Alias
		Companion string `json:"companion"`
	}{
		Alias:     (*Alias)(i),
		Companion: i.Affiliation.CompanionName(),
	})
}

func (i *Installment) InAdvance() bool {
	return i.DueDate.After(time.Now())
}
