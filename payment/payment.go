package payment

import (
	"oops-library-management/amount"
	"time"
)

type PaymentSystem interface {
	Compute(issuedDate time.Time) (amount.Amount, error)
}

type simplePaymentSystem struct {
}

func (s simplePaymentSystem) Compute(issuedDate time.Time) (amount.Amount, error) {
	now := time.Now()
	diff := now.Sub(issuedDate)
	days := diff.Hours() / 24
	if days < 8 {
		return amount.NewAmount(7)
	}
	daysLate := days - 7
	return amount.NewAmount(7 + (daysLate * 2))
}

func NewPaymentSystem() (PaymentSystem, error) {
	return simplePaymentSystem{}, nil
}
