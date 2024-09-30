package model

import "gorm.io/gorm"

type Chapter struct {
	gorm.Model
	Name                         string         `json:"name"`
	TreasurerRollingBalance      RollingBalance `json:"treasurer_rolling_balance"`
	TreasurerRollingBalanceID    uint           `json:"treasurer_rolling_balance_id"`
	GrandChapterRollingBalance   RollingBalance `json:"grand_chapter_rolling_balance"`
	GrandChapterRollingBalanceID uint           `json:"grand_chapter_rolling_balance_id"`
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
