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

func TestRollingBalance_Incomes(t *testing.T) {
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

	incomes := rb.Incomes()

	if len(incomes) != 1 {
		t.Errorf("expected 1 income, got %d", len(incomes))
	}

	if incomes[0] != movement1 {
		t.Errorf("expected income %v, got %v", movement1, incomes[0])
	}
}

func TestRollingBalance_TotalIncomes(t *testing.T) {
	rb := RollingBalance{}
	movement1 := &Movement{
		Amount: 100.0,
		Date:   time.Now(),
		Type:   CapitationPayment,
	}
	movement2 := &Movement{
		Amount: 50.0,
		Date:   time.Now(),
		Type:   CapitationPayment,
	}

	rb.AddMovement(movement1)
	rb.AddMovement(movement2)

	totalIncomes := rb.TotalIncomes()

	expectedTotal := 150.0
	if totalIncomes != expectedTotal {
		t.Errorf("expected total incomes %f, got %f", expectedTotal, totalIncomes)
	}
}

func TestRollingBalance_TotalExpenses(t *testing.T) {
	rb := RollingBalance{}
	movement1 := &Movement{
		Amount: 100.0,
		Date:   time.Now(),
		Type:   InstalmentCancellation,
	}
	movement2 := &Movement{
		Amount: 50.0,
		Date:   time.Now(),
		Type:   InstalmentCancellation,
	}

	rb.AddMovement(movement1)
	rb.AddMovement(movement2)

	totalExpenses := rb.TotalExpenses()

	expectedTotal := 150.0
	if totalExpenses != expectedTotal {
		t.Errorf("expected total expenses %f, got %f", expectedTotal, totalExpenses)
	}
}

func TestRollingBalance_Expenses(t *testing.T) {
	rb := RollingBalance{}
	movement1 := &Movement{
		Amount: 100.0,
		Date:   time.Now(),
		Type:   InstalmentCancellation,
	}
	movement2 := &Movement{
		Amount: 50.0,
		Date:   time.Now(),
		Type:   CapitationPayment,
	}

	rb.AddMovement(movement1)
	rb.AddMovement(movement2)

	expenses := rb.Expenses()

	if len(expenses) != 1 {
		t.Errorf("expected 1 expense, got %d", len(expenses))
	}

	if expenses[0] != movement1 {
		t.Errorf("expected expense %v, got %v", movement1, expenses[0])
	}
}
