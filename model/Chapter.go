package model

import (
	"gorm.io/gorm"
)

type Chapter struct {
	gorm.Model
	Name                      string          `json:"name" gorm:"unique"`
	TreasurerRollingBalance   *RollingBalance `json:"treasurer_rolling_balance"`
	TreasurerRollingBalanceID *uint           `json:"treasurer_rolling_balance_id"`
	Affiliations              []*Affiliation  `json:"affiliations" gorm:"foreignKey:ChapterID"`
	ChargeTypes               []*ChargeType   `json:"charge_types" gorm:"foreignKey:ChapterID"`
}

func (c *Chapter) AddMovement(movement *Movement) error {
	cc, err := GetInstalmentCancellation()
	if err != nil {
		return err
	}
	if movement.MovementType.Code == cc.Code {
		return nil
	}
	return c.TreasurerRollingBalance.AddMovement(movement)
}

func (c *Chapter) AddMovementTo(current float64, movement *Movement) float64 {
	if movement.Credit() {
		return current + movement.Amount
	}
	return current - movement.Amount
}

func (c *Chapter) Init() {
	c.TreasurerRollingBalance = &RollingBalance{}
	c.ChargeTypes = InitChargeTypes(c)
}

func (c *Chapter) PendingInstallments() []*Installment {
	out := make([]*Installment, 0)
	for _, affiliation := range c.Affiliations {
		for _, installment := range affiliation.PendingInstallments() {
			out = append(out, installment)
		}
	}
	return out
}

func (c *Chapter) DueInstallments() []*Installment {
	out := make([]*Installment, 0)
	for _, affiliation := range c.Affiliations {
		for _, installment := range affiliation.DueInstallments() {
			out = append(out, installment)
		}
	}
	return out
}

func (c *Chapter) PendingGreatChapterAmount() float64 {
	out := 0.0
	for _, installment := range c.PendingInstallments() {
		out += installment.GreatChapterAmount
	}
	return out
}

// AddAffiliation adds a new affiliation to the chapter and sets the charge
// of the brother in the chapter.For exaltation, the charge is ExaltationCharge.
func (c *Chapter) AddAffiliation(affiliation *Affiliation, charge *ChargeType) {
	c.Affiliations = append(c.Affiliations, affiliation)
	affiliation.SetChapter(c)
	affiliation.AddCharge(charge)
}

func (c *Chapter) AddDeposit(deposit *Deposit) error {
	m, err := deposit.CreateMovement()
	if err != nil {
		return err
	}
	err = c.TreasurerRollingBalance.AddMovement(m)
	if err != nil {
		return err
	}
	return nil
}

func (c *Chapter) UpdateInstallment(amount float64, greatChapterAmount float64) {
	for _, affiliation := range c.Affiliations {
		affiliation.UpdateInstallment(amount, greatChapterAmount)
	}
}

func (c *Chapter) PeriodPendingInstallments(brother *Brother) ([]*Installment, error) {
	period, err := GetCurrentPeriod()
	if err != nil {
		return nil, err
	}
	return period.PendingInstallments(brother, c)
}

func (c *Chapter) DueGreatChapterAmount() float64 {
	out := 0.0
	for _, affiliation := range c.Affiliations {
		out += affiliation.OverDueGreatChapter()
	}
	return out
}

func (c *Chapter) Deposits() []*Deposit {
	out := make([]*Deposit, 0)
	for _, affiliation := range c.Affiliations {
		for _, deposit := range affiliation.Deposits() {
			if !deposit.In(out) {
				out = append(out, deposit)
			}
		}
	}
	return out
}

func (c *Chapter) TotalDeposits() float64 {
	out := 0.0
	for _, deposit := range c.Deposits() {
		out += deposit.Amount
	}
	return out
}

func (c *Chapter) Balance() float64 {
	return c.TotalDeposits() - c.PendingGreatChapterAmount() - c.DueGreatChapterAmount()
}
