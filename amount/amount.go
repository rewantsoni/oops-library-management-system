package amount

import "errors"

type Amount struct {
	value float64
}

func NewAmount(value float64) (Amount, error) {
	if value < 1 {
		return Amount{}, errors.New("amount cannot be zero or negative")
	}
	return Amount{value: value}, nil
}
