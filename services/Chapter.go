package services

import (
	"argentina-tresury/db"
	"argentina-tresury/model"
)

func CreateChapter(chapter *model.Chapter) error {
	chapter.Init()
	if err := db.DB.Create(chapter).Error; err != nil {
		return err
	}
	return nil
}

func GetChapters() ([]model.Chapter, error) {
	var chapters []model.Chapter
	if err := db.DB.Find(&chapters).Error; err != nil {
		return nil, err
	}
	return chapters, nil
}

func GetChapter(u uint) (*model.Chapter, error) {
	var chapter model.Chapter
	if err := db.DB.First(&chapter, u).Error; err != nil {
		return nil, err
	}
	return &chapter, nil
}

func UpdateChapter(m *model.Chapter) error {
	if err := db.DB.Save(m).Error; err != nil {
		return err
	}
	return nil
}
