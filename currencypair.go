package money

import "regexp"

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

// Creates the currency pair from ISO string
// https://en.wikipedia.org/wiki/Currency_pair
// https://en.wikipedia.org/wiki/ISO_4217
func NewCurrencyPairFromIso(iso string) (*CurrencyPair, error) {
	regex := regexp.MustCompile("^([A-Z]{2,3})/([A-Z]{2,3}) ([0-9]*\\.?[0-9]+)$")

	if !regex.Match([]byte(iso)) {
		return nil, ErrInvalidIsoPair
	}

	regex.FindAllString(iso, -1)

	return nil, nil
}

// Converts from base to counter currency
func (p CurrencyPair) Convert(money *Money) (*Money, error) {
	if !money.Currency.Equals(p.Base) {
		return nil, ErrNotSameCurrency
	}

	return money.Convert(p.Counter, p.Ratio), nil
}
