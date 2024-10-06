package model

import (
	"gorm.io/gorm"
	"time"
)

type Chapter struct {
	gorm.Model
	Name                         string          `json:"name"`
	TreasurerRollingBalance      *RollingBalance `json:"treasurer_rolling_balance"`
	TreasurerRollingBalanceID    *uint           `json:"treasurer_rolling_balance_id"`
	GrandChapterRollingBalance   *RollingBalance `json:"grand_chapter_rolling_balance"`
	GrandChapterRollingBalanceID *uint           `json:"grand_chapter_rolling_balance_id"`
	Affiliations                 []*Affiliation  `json:"affiliations" gorm:"foreignKey:ChapterID"`
	CurrentPeriodID              *uint           `json:"current_period_id"`
	CurrentPeriod                *Period         `json:"current_period"`
}

func (c *Chapter) AddMovement(movement *Movement) {
	c.TreasurerRollingBalance.AddMovement(movement)
}

func (c *Chapter) AddMovementTo(current float64, movement *Movement) float64 {
	if movement.Credit() {
		return current + movement.Amount
	}
	return current - movement.Amount
}

func (c *Chapter) Init() {
	c.GrandChapterRollingBalance = &RollingBalance{}
	c.TreasurerRollingBalance = &RollingBalance{}
}

func (c *Chapter) PendingInstallments() []*Installment {
	out := make([]*Installment, 0)
	for _, affiliation := range c.Affiliations {
		out = append(out, affiliation.PendingInstallments()...)
	}
	return out
}

func (c *Chapter) PendingGrandChapterAmount() float64 {
	out := 0.0
	for _, installment := range c.PendingInstallments() {
		out += installment.GrandChapterAmount
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
	mg, err := deposit.GrandChapterMovement()
	if err != nil {
		return err
	}
	c.TreasurerRollingBalance.AddMovement(m)
	c.GrandChapterRollingBalance.AddMovement(mg)
	return nil
}

func (c *Chapter) AddExpenses(amount float64, date time.Time, expenseType *MovementType, description string) {
	c.TreasurerRollingBalance.AddMovement(&Movement{
		Amount:       amount,
		MovementType: expenseType,
		Description:  description,
		Date:         date,
	})

}

func (c *Chapter) AddBrotherExpense(amount float64, date time.Time, brother *Brother, expenseType *MovementType,
	description string) {
	affiliation := c.AffiliationOf(brother)
	mov := Movement{
		Amount:       amount,
		MovementType: expenseType,
		Description:  description + " pagado por Hermano" + brother.FirstName + " " + brother.LastNames,
		Date:         date,
	}
	if affiliation != nil {
		affiliation.AddMovement(&mov)
	} else {
		c.AddMovement(&mov)
	}

}

func (c *Chapter) AddBrotherMovement(brother *Brother, movement *Movement) {
	affiliation := c.AffiliationOf(brother)
	if affiliation != nil {
		affiliation.AddMovement(movement)
	} else {
		c.AddMovement(movement)
	}
}

func (c *Chapter) UpdateInstallment(amount float64, greatChapterAmount float64, movement *Movement) {
	amount = 0.0
	for _, affiliation := range c.Affiliations {
		amount += affiliation.UpdateInstallment(amount, greatChapterAmount)
	}
	c.GrandChapterRollingBalance.AddMovement(movement)
}

func (c *Chapter) GenerateGrandChapterMonthlyDebit(month uint) {
	//TODO: Implement
}

func (c *Chapter) AffiliationOf(brother *Brother) *Affiliation {
	for _, affiliation := range c.Affiliations {
		if *affiliation.BrotherID == brother.ID {
			return affiliation
		}
	}
	return nil
}

func (c *Chapter) PeriodPendingInstallments(brother *Brother) []*Installment {
	return c.CurrentPeriod.PendingInstallments(brother)
}
