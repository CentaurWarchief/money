package money_test

import (
	"testing"

	"github.com/CentaurWarchief/money"
	"github.com/stretchr/testify/assert"
)

func TestAllocate(t *testing.T) {
	m := money.NewMoney(100, money.EUR)

	p := m.Allocate([]float64{
		1,
		1,
		1,
	})

	assert.Equal(t, int64(34), p[0].Amount)
	assert.Equal(t, int64(33), p[1].Amount)
	assert.Equal(t, int64(33), p[2].Amount)

	m = money.NewMoney(101, money.EUR)

	p = m.Allocate([]float64{
		1,
		1,
		1,
	})

	assert.Equal(t, int64(34), p[0].Amount)
	assert.Equal(t, int64(34), p[1].Amount)
	assert.Equal(t, int64(33), p[2].Amount)
}

func TestAllocationOrder(t *testing.T) {
	m := money.NewMoney(5, money.EUR)

	p := m.Allocate([]float64{
		3,
		7,
	})

	assert.Equal(t, int64(2), p[0].Amount)
	assert.Equal(t, int64(3), p[1].Amount)

	p = m.Allocate([]float64{
		7,
		3,
	})

	assert.Equal(t, int64(4), p[0].Amount)
	assert.Equal(t, int64(1), p[1].Amount)
}
