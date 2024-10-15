package model

import (
	"gorm.io/gorm"
	"time"
)

type Period struct {
	gorm.Model
	Year                  int  `json:"year"`
	TotalInstallments     int  `json:"total_installments"`
	FirstMonthInstallment int  `json:"first_month_installment"`
	Current               bool `json:"current"`
}

// PendingInstallments returns the installments that are pending for this period
// for the companion.This method has the assumption what the installments are monthly
// and the due date is the 10th of each month.
func (p *Period) PendingInstallments(companion *Companion, chapter *Chapter) ([]*Installment, error) {
	out := make([]*Installment, 0)
	ct, err := GetMonthlyCharge(chapter)
	if err != nil {
		return nil, err
	}
	for i := 1; i <= p.TotalInstallments+p.FirstMonthInstallment-1; i++ {
		if i >= p.FirstMonthInstallment && i >= int(time.Now().Month()) {
			installment := &Installment{
				Year:               p.Year,
				Month:              i,
				Amount:             companion.InstallmentAmount(ct),
				GreatChapterAmount: companion.GreatChapterAmount(ct),
				Paid:               false,
				OnTheSpot:          false,
				//TODO: I18n
				Description: "Cuota mensual " + time.Month(i).String(),
				DueDate:     time.Date(p.Year, time.Month(i), 10, 0, 0, 0, 0, time.Local),
				ChargeType:  ct,
			}
			out = append(out, installment)
		}
	}
	return out, nil
}

func GetCurrentPeriod() (*Period, error) {
	period := &Period{Current: true}
	if err := DB.Model(period).First(period).Error; err != nil {
		return nil, err
	}
	return period, nil
}
