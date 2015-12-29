package money

import (
	"encoding/json"
	"math"
)

// NewMoney creates a new money representation of the currency with given amount
func NewMoney(amount int64, currency Currency) *Money {
	return &Money{
		amount,
		currency,
	}
}

// Money represents an amount of money in the given currency
type Money struct {
	Amount   int64
	Currency Currency
}

// IsSameCurrency compares the currency equality of both money representations
func (m Money) IsSameCurrency(money Money) bool {
	return m.Currency.Equals(money.Currency)
}

// Compare compares the money amount (-1 if less, 0 if equals and 1 if greater)
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

// IsGreaterThan returns whether the money amount is greater
func (m Money) IsGreaterThan(money Money) bool {
	return 1 == m.Compare(money)
}

// IsGreaterThanOrEqual returns whether the money amount is greater than or equals
func (m Money) IsGreaterThanOrEqual(money Money) bool {
	return 0 <= m.Compare(money)
}

// IsLessThan returns whether the money amount is less
func (m Money) IsLessThan(money Money) bool {
	return -1 == m.Compare(money)
}

// Equals returns true if both money currency and amount are equals
func (m Money) Equals(money Money) bool {
	return m.IsSameCurrency(money) && m.Amount == money.Amount
}

// IsZero checks whether the current money amount is zero
func (m Money) IsZero() bool {
	return m.Amount == int64(0)
}

// IsPositive checks whether the current money amount is above 0
func (m Money) IsPositive() bool {
	return m.Amount > int64(0)
}

// IsNegative checks whether the current money amount is below 0
func (m Money) IsNegative() bool {
	return m.Amount < int64(0)
}

// Convert converts the current amount of money to the target currency
// using the given rate
func (m Money) Convert(target Currency, rate float64) *Money {
	return NewMoney(int64((float64(m.Amount) * rate)), target)
}

// Allocate distributes the money amount by the given ratios
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

	for i := int64(0); i < r; i++ {
		results[i].Amount++
	}

	return results
}

// AllocateTo allocates the money among the number of targets
func (m Money) AllocateTo(targets int) []*Money {
	amount := int64(m.Amount / int64(targets))

	results := make([]*Money, targets)

	for i := 0; i < targets; i++ {
		results[i] = NewMoney(amount, m.Currency)
	}

	for i := int64(0); i < (m.Amount % int64(targets)); i++ {
		results[i].Amount++
	}

	return results
}

// MarshalJSON marshals the current Money representation to JSON
func (m Money) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"amount":   m.Amount,
		"currency": m.Currency,
	})
}
