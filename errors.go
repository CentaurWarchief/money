package money

import "errors"

var (
	// ErrNotSameCurrency is the error returned when an operation
	// require both currencies to be equal
	ErrNotSameCurrency = errors.New("different currencies")
	// ErrInvalidIsoPair is the error returned when the provided ISO string
	// for creating a CurrencyPair does not match the required pattern
	ErrInvalidIsoPair = errors.New("invalid currency pair ISO string")
)
