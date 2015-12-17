package money_test

import (
	"testing"

	"github.com/CentaurWarchief/money"
	"github.com/stretchr/testify/assert"
)

func TestCurrencyEquals(t *testing.T) {
	assert.True(t, money.EUR.Equals(money.EUR))
}

func TestCurrencyNotEquals(t *testing.T) {
	eur := money.EUR
	gbp := money.GBP

	assert.False(t, eur.Equals(gbp))
}
