package money

// Currency is a type for representing a currency
type Currency string

// Equals compares this currency against another currency
func (c Currency) Equals(other Currency) bool {
	return c == other
}
