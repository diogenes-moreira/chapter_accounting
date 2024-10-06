package services

import (
	"argentina-tresury/model"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDBChapter() {
	model.DB, _ = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	err := model.DB.AutoMigrate(&model.Chapter{})
	if err != nil {
		return
	}
}

func TestCreateChapter(t *testing.T) {
	setupTestDBChapter()
	chapter := &model.Chapter{Name: "Sample Chapter"}

	err := CreateChapter(chapter)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var result model.Chapter
	model.DB.First(&result, chapter.ID)
	if result.Name != "Sample Chapter" {
		t.Errorf("expected name to be 'Sample Chapter', got %s", result.Name)
	}
}

func TestGetChapters(t *testing.T) {
	setupTestDBChapter()
	chapter1 := &model.Chapter{Name: "Chapter One"}
	chapter2 := &model.Chapter{Name: "Chapter Two"}
	model.DB.Create(chapter1)
	model.DB.Create(chapter2)

	chapters, err := GetChapters()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(chapters) != 2 {
		t.Errorf("expected 2 chapters, got %d", len(chapters))
	}
}

func TestGetChapter(t *testing.T) {
	setupTestDBChapter()
	chapter := &model.Chapter{Name: "Sample Chapter"}
	model.DB.Create(chapter)

	result, err := GetChapter(chapter.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Name != "Sample Chapter" {
		t.Errorf("expected name to be 'Sample Chapter', got %s", result.Name)
	}
}

func TestUpdateChapter(t *testing.T) {
	setupTestDBChapter()
	chapter := &model.Chapter{Name: "Sample Chapter"}
	model.DB.Create(chapter)

	chapter.Name = "Updated Chapter"
	err := UpdateChapter(chapter)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var result model.Chapter
	model.DB.First(&result, chapter.ID)
	if result.Name != "Updated Chapter" {
		t.Errorf("expected name to be 'Updated Chapter', got %s", result.Name)
	}
}
