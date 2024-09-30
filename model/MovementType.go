package model

var movementTypes map[string]bool

const CapitationPayment = "capitation_payment"
const InstalmentCancellation = "instalment_cancellation"

func IsCredit(name string) bool {
	if movementTypes == nil {
		movementTypes := make(map[string]bool)
		movementTypes[CapitationPayment] = true
		movementTypes[InstalmentCancellation] = false
	}
	return movementTypes[name]
}
