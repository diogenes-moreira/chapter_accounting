package model

import (
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDBRollingBalance() {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&RollingBalance{}, &Movement{})
}

func TestAddMovement(t *testing.T) {
	setupTestDBRollingBalance()
	rb := &RollingBalance{}
	mt := &MovementType{Code: "Test Movement Type"}
	mov := &Movement{
		Amount:       100.0,
		MovementType: mt,
		Description:  "Test Movement",
		Date:         time.Now(),
	}

	rb.AddMovement(mov)

	if len(rb.Movements) != 1 {
		t.Errorf("expected 1 movement, got %d", len(rb.Movements))
	}

	if rb.Movements[0].Amount != 100.0 {
		t.Errorf("expected amount to be 100.0, got %f", rb.Movements[0].Amount)
	}
}

func TestGetBalance(t *testing.T) {
	setupTestDBRollingBalance()
	rb := &RollingBalance{}
	mt := &MovementType{Code: "Test Movement Type"}
	mov1 := &Movement{
		Amount:       100.0,
		MovementType: mt,
		Description:  "Test Movement 1",
		Date:         time.Now(),
	}
	mov2 := &Movement{
		Amount:       -50.0,
		MovementType: mt,
		Description:  "Test Movement 2",
		Date:         time.Now(),
	}

	rb.AddMovement(mov1)
	rb.AddMovement(mov2)

	balance := rb.Balance(&Affiliation{})
	expectedBalance := 50.0

	if balance != expectedBalance {
		t.Errorf("expected balance to be %f, got %f", expectedBalance, balance)
	}
}
