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

// TODO: Implementar el método AddAffiliation
// TODO: Implementar el método AddExaltation
// TODO: Implementar el método AddDeposit

func (c *Chapter) AddExpenses(amount float64, date time.Time, expenseType string) {
	c.TreasurerRollingBalance.AddMovement(&Movement{
		Amount:      amount,
		Type:        expenseType,
		Description: "Gasto " + expenseType,
		Date:        date,
	})

}

func (c *Chapter) AddBrotherExpense(amount float64, date time.Time, brother *Brother, expenseType string) {
	affiliation := c.AffiliationOf(brother)
	mov := Movement{
		Amount:      amount,
		Type:        expenseType,
		Description: "Gasto pagado por Hermano" + brother.Name,
		Date:        date,
	}
	if affiliation != nil {
		affiliation.AddMovement(&mov)
	} else {
		c.AddMovement(&mov)
	}

}

func (c *Chapter) AddBrotherIncome(amount float64, date time.Time, brother *Brother) {
	affiliation := c.AffiliationOf(brother)
	mov := Movement{
		Amount:      amount,
		Type:        MovementTypeBrotherIncome,
		Description: "Ingreso de Hermano" + brother.Name,
		Date:        date,
	}
	if affiliation != nil {
		affiliation.AddMovement(&mov)
	} else {
		c.AddMovement(&mov)
	}

}

func (c *Chapter) AddBagIncome(amount float64, date time.Time) {
	c.TreasurerRollingBalance.AddMovement(&Movement{
		Amount:      amount,
		Type:        MovementTypeBagIncome,
		Description: "Ingreso de Saco",
		Date:        date,
	})
}

func (c *Chapter) UpdateInstallment(amount float64, greatChapterAmount float64) {
	amount = 0.0
	for _, affiliation := range c.Affiliations {
		amount += affiliation.UpdateInstallment(amount, greatChapterAmount)
	}
	c.GrandChapterRollingBalance.AddMovement(&Movement{
		Amount:  amount,
		Type:    AdjustmentOfInstallments,
		Receipt: "",
		Date:    time.Time{},
	})
}

func (c *Chapter) GenerateGrandChapterMonthlyDebit(month uint) {
	amount := 0.0
	for _, affiliation := range c.Affiliations {
		amount += affiliation.GrandChapterAmountDueAt(month)
	}
	c.GrandChapterRollingBalance.AddMovement(&Movement{
		Amount:  amount,
		Type:    MovementTypeGrandChapterMonthlyDebit,
		Receipt: "",
		Date:    time.Time{},
	})
}

func (c *Chapter) AffiliationOf(brother *Brother) *Affiliation {
	for _, affiliation := range c.Affiliations {
		if *affiliation.BrotherID == brother.ID {
			return affiliation
		}
	}
	return nil
}
