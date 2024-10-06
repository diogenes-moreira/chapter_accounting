package model

import (
	"gorm.io/gorm"
	"time"
)

type Period struct {
	gorm.Model
	Year                  int `json:"year"`
	TotalInstallments     int `json:"total_installments"`
	FirstMonthInstallment int `json:"first_month_installment"`
}

func (p *Period) IsCurrent() bool {
	return p.Year == time.Now().Year()
}

// PendingInstallments returns the installments that are pending for this period
// for the brother.This method has the assumption what the installments are monthly
// and the due date is the 10th of each month.
func (p *Period) PendingInstallments(brother *Brother) []*Installment {
	out := make([]*Installment, 0)
	for i := 1; i <= p.TotalInstallments; i++ {
		if i >= p.FirstMonthInstallment {
			installment := &Installment{
				Year:               p.Year,
				Month:              i,
				Amount:             brother.InstallmentAmount(),
				GrandChapterAmount: brother.GreatChapterAmount(),
				Paid:               false,
				OnTheSpot:          false,
				//TODO: I18n
				Description: "Cuota mensual " + time.Month(i).String(),
				DueDate:     time.Date(p.Year, time.Month(i), 10, 0, 0, 0, 0, time.Local),
			}
			out = append(out, installment)
		}
	}
	return out
}
