package money

type Currency string

// Compare this currency against another currency
func (c Currency) Equals(other Currency) bool {
	return c == other
}
