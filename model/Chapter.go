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
	Affiliations                 []Affiliation   `json:"affiliations" gorm:"foreignKey:ChapterID"`
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
// TODO: Implementar el método AddExpense
// TODO: Implementar el método AddBrotherExpense
// TODO: Implementar el método AddBrotherIncome
// TODO: Implementar el método AddBagIncome
// TODO: Implementar el método UpdateInstallment
func (c *Chapter) UpdateInstallment(amount float64, greatChapterAmount float64) {
	amount = 0.0
	for _, affiliation := range c.Affiliations {
		amount += affiliation.UpdateInstallment(amount, greatChapterAmount)
	}

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
