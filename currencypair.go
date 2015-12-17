package money

type CurrencyPair struct {
	Base    Currency
	Counter Currency
	Ratio   float64
}

// Creates a new currency pair based on the given ratio
func NewCurrencyPair(base, counter Currency, ratio float64) *CurrencyPair {
	return &CurrencyPair{
		base,
		counter,
		ratio,
	}
}

// Converts the given amount of money based on the ratio between both currencies
func (p CurrencyPair) Convert(money *Money) (*Money, error) {
	if !money.Currency.Equals(p.Base) {
		return nil, ErrNotSameCurrency
	}

	return NewMoney(int64((float64(money.Amount) * p.Ratio)), p.Counter), nil
}
