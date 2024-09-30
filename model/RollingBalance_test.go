package model

import (
	"testing"
	"time"
)

func TestRollingBalance_AddMovement(t *testing.T) {
	rb := RollingBalance{}
	movement := &Movement{
		Amount: 100.0,
		Date:   time.Now(),
		Type:   CapitationPayment,
	}

	rb.AddMovement(movement)

	if len(rb.Movements) != 1 {
		t.Errorf("expected 1 movement, got %d", len(rb.Movements))
	}

	if rb.Movements[0] != movement {
		t.Errorf("expected movement %v, got %v", movement, rb.Movements[0])
	}
}

func TestRollingBalance_Balance(t *testing.T) {
	rb := RollingBalance{}
	movement1 := &Movement{
		Amount: 100.0,
		Date:   time.Now(),
		Type:   CapitationPayment,
	}
	movement2 := &Movement{
		Amount: 50.0,
		Date:   time.Now(),
		Type:   InstalmentCancellation,
	}

	rb.AddMovement(movement1)
	rb.AddMovement(movement2)

	affiliation := &Affiliation{}
	balance := rb.Balance(affiliation)

	expectedBalance := 50.0
	if balance != expectedBalance {
		t.Errorf("expected balance %f, got %f", expectedBalance, balance)
	}
}
