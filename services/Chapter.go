package services

import (
	"argentina-tresury/model"
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
		Preload(`TreasurerRollingBalance.Movements`).
		Preload("TreasurerRollingBalance.Movements.MovementType").First(&chapter, u).Error; err != nil {
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

func UpdateChapter(m *model.Chapter) error {
	if err := model.DB.Save(m).Error; err != nil {
		return err
	}
	return nil
}

func CreateAffiliation(brother *model.Brother, chapter *model.Chapter, isHonorary bool) (*model.Affiliation, error) {
	affiliation := &model.Affiliation{
		Period:         chapter.CurrentPeriod,
		Brother:        brother,
		Installments:   []*model.Installment{},
		StartDate:      time.Now(),
		EndDate:        nil,
		RollingBalance: model.RollingBalance{},
		Balance:        0,
		Honorary:       isHonorary,
	}

	if err := model.DB.Create(affiliation).Error; err != nil {
		return nil, err
	}
	affiliation.SetChapter(chapter)
	return affiliation, nil
}
