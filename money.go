package money

import "math"

type Money struct {
	Amount   int64
	Currency Currency
}

// Creates a money representation of the currency with given amount
func NewMoney(amount int64, currency Currency) *Money {
	return &Money{
		amount,
		currency,
	}
}

// Compares the currency equality of both money representations
func (m Money) IsSameCurrency(money Money) bool {
	return m.Currency.Equals(money.Currency)
}

// Compares the money amount (-1 if less, 0 if equals and 1 if greater)
func (m Money) Compare(money Money) int {
	if !m.IsSameCurrency(money) {
		panic(ErrNotSameCurrency)
	}

	if m.Amount < money.Amount {
		return -1
	}

	if m.Amount == money.Amount {
		return 0
	}

	return 1
}

// Returns whether the money amount is greater
func (m Money) IsGreaterThan(money Money) bool {
	return 1 == m.Compare(money)
}

// Returns whether the money amount is greater than or equals
func (m Money) IsGreaterThanOrEqual(money Money) bool {
	return 0 <= m.Compare(money)
}

// Returns whether the money amount is less
func (m Money) IsLessThan(money Money) bool {
	return -1 == m.Compare(money)
}

// Returns true if both money currency and amount are equals
func (m Money) Equals(money Money) bool {
	return m.IsSameCurrency(money) && m.Amount == money.Amount
}

// Checks whether the current money amount is zero
func (m Money) IsZero() bool {
	return m.Amount == int64(0)
}

// Checks whether the current money amount is above 0
func (m Money) IsPositive() bool {
	return m.Amount > int64(0)
}

// Checks whether the current money amount is below 0
func (m Money) IsNegative() bool {
	return m.Amount < int64(0)
}

// Adds a specific amount of money
func (m Money) Add(money Money) (*Money, error) {
	if !m.IsSameCurrency(money) {
		return nil, ErrNotSameCurrency
	}

	return NewMoney((m.Amount + money.Amount), m.Currency), nil
}

// Subtracts a specific amount of money
func (m Money) Subtract(money Money) (*Money, error) {
	if !m.IsSameCurrency(money) {
		return nil, ErrNotSameCurrency
	}

	return NewMoney((m.Amount - money.Amount), m.Currency), nil
}

// Multiplies the money amount
func (m Money) Multiply(money Money) (*Money, error) {
	if !m.IsSameCurrency(money) {
		return nil, ErrNotSameCurrency
	}

	return NewMoney(int64((m.Amount * money.Amount)), m.Currency), nil
}

// Divides the money amount
func (m Money) Divide(money Money) (*Money, error) {
	if !m.IsSameCurrency(money) {
		return nil, ErrNotSameCurrency
	}

	return NewMoney(int64((m.Amount / money.Amount)), m.Currency), nil
}

// Distributes the money amount by the given ratios
func (m Money) Allocate(ratios []float64) (results []*Money) {
	r := m.Amount
	t := 0.0

	for _, ratio := range ratios {
		t += ratio
	}

	for _, ratio := range ratios {
		share := math.Floor(float64(m.Amount) * ratio / t)
		r -= int64(share)

		results = append(
			results,
			NewMoney(int64(share), m.Currency),
		)
	}

	for i := 0; r > 0; i++ {
		results[i].Amount++
		r--
	}

	return results
}
