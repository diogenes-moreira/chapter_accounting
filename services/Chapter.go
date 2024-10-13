package services

import (
	"argentina-tresury/model"
	"gorm.io/gorm"
	"time"
)

func CreateChapter(chapter *model.Chapter) error {
	chapter.Init()
	if err := model.DB.Create(chapter).Error; err != nil {
		return err
	}
	return nil
}

func GetChapters() ([]model.Chapter, error) {
	var chapters []model.Chapter
	if err := model.DB.Find(&chapters).Error; err != nil {
		return nil, err
	}
	return chapters, nil
}

func GetChapter(u uint) (*model.Chapter, error) {
	var chapter model.Chapter
	if err := model.DB.
		Preload(`TreasurerRollingBalance`).
		Preload(`TreasurerRollingBalance.Movements`, func(db *gorm.DB) *gorm.DB {
			return db.Order("Movements.Date")
		}).
		Preload("TreasurerRollingBalance.Movements.MovementType").
		Preload("ChargeTypes").
		First(&chapter, u).Error; err != nil {
		return nil, err
	}
	chapter.TreasurerRollingBalance.Adder = &chapter
	return &chapter, nil
}

func GetChapterAffiliations(u uint) ([]*model.Affiliation, error) {
	var chapter model.Chapter
	var out []*model.Affiliation
	if err := model.DB.Preload("Affiliations").
		Preload("Affiliations.Brother").
		Preload("Affiliations.RollingBalance").
		Preload("Affiliations.RollingBalance.Movements", func(db *gorm.DB) *gorm.DB {
			return db.Order("Movements.ID")
		}).
		Preload("Affiliations.RollingBalance.Movements.MovementType").
		Preload("Affiliations.Installments").
		Preload("Affiliations.Period").First(&chapter, u).Error; err != nil {
		return nil, err
	}
	for _, affiliation := range chapter.Affiliations {
		if affiliation.IsCurrent() {
			out = append(out, affiliation)
		}
	}
	return out, nil
}

func CreateChapterMovement(chapter uint, amount float64, receipt string, date time.Time, typeMovement string,
	description string) error {
	mt, err := model.GetMovementType(typeMovement)
	if err != nil {
		return err
	}
	c, err := GetChapter(chapter)
	if err != nil {
		return err
	}
	movement := &model.Movement{
		Amount:       amount,
		Receipt:      receipt,
		Date:         date,
		Description:  description,
		MovementType: mt,
	}
	return c.AddMovement(movement)
}
func UpdateChapter(m *model.Chapter) error {
	if err := model.DB.Save(m).Error; err != nil {
		return err
	}
	return nil
}

func CreateAffiliation(brother *model.Brother, chapter *model.Chapter, isHonorary bool) (*model.Affiliation, error) {
	rollingBalance := model.RollingBalance{}
	if err := model.DB.Create(&rollingBalance).Error; err != nil {
		return nil, err
	}
	period, err := model.GetCurrentPeriod()
	if err != nil {
		return nil, err
	}
	affiliation := &model.Affiliation{
		Period:         period,
		Brother:        brother,
		Installments:   []*model.Installment{},
		StartDate:      time.Now(),
		EndDate:        nil,
		RollingBalance: rollingBalance,
		Balance:        0,
		Honorary:       isHonorary,
	}

	if err := model.DB.Create(affiliation).Error; err != nil {
		return nil, err
	}
	err = affiliation.SetChapter(chapter)
	if err != nil {
		return nil, err
	}
	if err := model.DB.Save(affiliation).Error; err != nil {
		return nil, err
	}
	return affiliation, nil
}
