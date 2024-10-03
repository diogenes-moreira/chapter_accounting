package model

var movementTypes map[string]bool

const CapitationPayment = "capitation_payment"
const InstalmentCancellation = "instalment_cancellation"
const LibraryExpenses = "library_expenses"
const MovementTypeGrandChapterMonthlyDebit = "grand_chapter_monthly_debit"
const AdjustmentOfInstallments = "adjustment_of_installments"
const MovementTypeBagIncome = "bag_income"
const MovementTypeBrotherIncome = "brother_income"

func IsCredit(name string) bool {
	if movementTypes == nil {
		movementTypes = make(map[string]bool)
		movementTypes[CapitationPayment] = true
		movementTypes[InstalmentCancellation] = false
		movementTypes[LibraryExpenses] = false
		movementTypes[MovementTypeGrandChapterMonthlyDebit] = false
		movementTypes[AdjustmentOfInstallments] = false
		movementTypes[MovementTypeBagIncome] = true
		movementTypes[MovementTypeBrotherIncome] = true
	}
	return movementTypes[name]
}
