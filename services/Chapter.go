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
		Preload("Affiliations").
		Preload("Affiliations.Brother").
		Preload("Affiliations.Installments").
		Preload("Affiliations.Installments.Affiliation").
		Preload("Affiliations.Installments.Affiliation.Brother").
		Preload("Affiliations.Installments.Deposit").
		Preload("Affiliations.Installments.Deposit.Installments").
		Preload("Affiliations.Installments.Deposit.Installments.Affiliation.Brother").
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
		Preload("Affiliations.Installments.Deposit").
		Preload("Affiliations.Installments.Affiliation").
		Preload("Affiliations.Installments.Affiliation.Brother").
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
	err = c.AddMovement(movement)
	if err != nil {
		return err
	}
	return model.DB.Save(c).Error
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

type GreatChapterStatus struct {
	PendingInstallments []*model.Installment `json:"pending_installments"`
	PendingAmount       float64              `json:"pending_amount"`
	DueInstallments     []*model.Installment `json:"due_installments"`
	Deposits            []*model.Deposit     `json:"deposits"`
	TotalDeposits       float64              `json:"total_deposits"`
	DueAmount           float64              `json:"due_amount"`
	Balance             float64              `json:"balance"`
}

func CreateDeposit(chapterId uint, installmentsId []uint64) error {
	chapter, err := GetChapter(chapterId)
	var installments []*model.Installment
	if err != nil {
		return err
	}
	for _, installmentId := range installmentsId {
		i := &model.Installment{}
		err := model.DB.First(i, installmentId).Error
		if err != nil {
			return err
		}
		installments = append(installments, i)
	}
	deposit := &model.Deposit{
		Amount:       0,
		Installments: []*model.Installment{},
		DepositDate:  time.Now(),
	}
	deposit.AddInstallments(installments)
	err = model.DB.Create(deposit).Error
	if err != nil {
		return err
	}
	movement, err := deposit.CreateMovement()
	if err != nil {
		return err
	}
	err = chapter.AddMovement(movement)
	if err != nil {
		return err
	}
	err = model.DB.Save(chapter).Error
	for _, installment := range installments {
		err = model.DB.Save(installment).Error
	}
	if err != nil {
		return err
	}
	return nil
}

func GetGreatChapterStatus(chapterId uint) (*GreatChapterStatus, error) {
	chapter, err := GetChapter(chapterId)
	if err != nil {
		return nil, err
	}
	if chapter == nil {
		return nil, nil
	}
	gcs := &GreatChapterStatus{
		PendingInstallments: chapter.PendingInstallments(),
		PendingAmount:       chapter.PendingGreatChapterAmount(),
		DueInstallments:     chapter.DueInstallments(),
		DueAmount:           chapter.DueGreatChapterAmount(),
		Deposits:            chapter.Deposits(),
		TotalDeposits:       chapter.TotalDeposits(),
		Balance:             chapter.Balance(),
	}
	return gcs, nil
}
