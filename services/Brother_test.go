package services

import (
	"argentina-tresury/model"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() {
	model.DB, _ = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	err := model.DB.AutoMigrate(&model.Brother{})
	if err != nil {
		return
	}
}

func TestCreateBrother(t *testing.T) {
	setupTestDB()
	brother := &model.Brother{Name: "John Doe"}

	err := CreateBrother(brother)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var result model.Brother
	model.DB.First(&result, brother.ID)
	if result.Name != "John Doe" {
		t.Errorf("expected name to be 'John Doe', got %s", result.Name)
	}
}

func TestGetBrothers(t *testing.T) {
	setupTestDB()
	brother1 := &model.Brother{Name: "John Doe"}
	brother2 := &model.Brother{Name: "Jane Doe"}
	model.DB.Create(brother1)
	model.DB.Create(brother2)

	brothers, err := GetBrothers()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(brothers) != 2 {
		t.Errorf("expected 2 brothers, got %d", len(brothers))
	}
}

func TestGetBrother(t *testing.T) {
	setupTestDB()
	brother := &model.Brother{Name: "John Doe"}
	model.DB.Create(brother)

	result, err := GetBrother(brother.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Name != "John Doe" {
		t.Errorf("expected name to be 'John Doe', got %s", result.Name)
	}
}

func TestUpdateBrother(t *testing.T) {
	setupTestDB()
	brother := &model.Brother{Name: "John Doe"}
	model.DB.Create(brother)

	brother.Name = "John Smith"
	err := UpdateBrother(brother)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var result model.Brother
	model.DB.First(&result, brother.ID)
	if result.Name != "John Smith" {
		t.Errorf("expected name to be 'John Smith', got %s", result.Name)
	}
}
