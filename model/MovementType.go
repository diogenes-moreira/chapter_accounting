package model

var movementTypes map[string]bool

const CapitationPayment = "capitation_payment"
const InstalmentCancellation = "instalment_cancellation"
const LibraryExpenses = "library_expenses"
const MovementTypeGrandChapterMonthlyDebit = "grand_chapter_monthly_debit"

func IsCredit(name string) bool {
	if movementTypes == nil {
		movementTypes = make(map[string]bool)
		movementTypes[CapitationPayment] = true
		movementTypes[InstalmentCancellation] = false
		movementTypes[LibraryExpenses] = false
		movementTypes[MovementTypeGrandChapterMonthlyDebit] = false
	}
	return movementTypes[name]
}
