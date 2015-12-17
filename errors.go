package money

import "errors"

var (
	ErrNotSameCurrency = errors.New("different currencies")
	ErrInvalidIsoPair  = errors.New("invalid currency pair ISO string")
)
