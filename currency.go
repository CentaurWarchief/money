package money

type Currency string

// Compare this currency against another currency
func (c Currency) Equals(candidate Currency) bool {
	return c == candidate
}
