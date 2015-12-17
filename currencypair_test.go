package money_test

import (
	"testing"

	"github.com/CentaurWarchief/money"
	"github.com/stretchr/testify/assert"
)

func TestConvertDifferentCurrencies(t *testing.T) {
	pair := money.NewCurrencyPair(money.EUR, money.USD, 1.25)

	gbp, err := pair.Convert(money.NewMoney(100, money.GBP))

	assert.NotNil(t, err)
	assert.Nil(t, gbp)
	assert.Equal(t, money.ErrNotSameCurrency, err)
}

func TestConvertEurToUsd(t *testing.T) {
	eur := money.NewMoney(100, money.EUR)

	var pair *money.CurrencyPair

	pair = money.NewCurrencyPair(money.EUR, money.USD, 1.25)

	usd, err := pair.Convert(eur)

	assert.Nil(t, err)
	assert.Equal(t, int64(125), usd.Amount)
	assert.Equal(t, usd.Currency, money.USD)

	pair = money.NewCurrencyPair(money.USD, money.EUR, 0.80)

	eur, err = pair.Convert(usd)

	assert.Nil(t, err)
	assert.Equal(t, int64(100), eur.Amount)
	assert.Equal(t, eur.Currency, money.EUR)
}

func TestNewCurrencyPairFromInvalidIsoString(t *testing.T) {
	for _, iso := range []string{
		"",
		"/",
		"EUR/USD *****",
		"EUR/USD -",
		"EUR/USD 1.25#",
		"EUR/USD",
		"EUR",
	} {
		pair, err := money.NewCurrencyPairFromIso(iso)

		assert.Nil(t, pair)
		assert.Equal(t, money.ErrInvalidIsoPair, err, iso)
	}
}
