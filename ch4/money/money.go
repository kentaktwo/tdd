package money

// Dollar is a struct to express dollar.
type Dollar struct {
	amount int
}

// NewDollar is a constructor of Dollar.
func NewDollar(amount int) *Dollar {
	dollar := &Dollar{amount: amount}
	return dollar
}

// Times multiplies number to Dollar.
func (d *Dollar) Times(multiplier int) *Dollar {
	return NewDollar(d.amount * multiplier)
}

// Equals compares two Dollars.
func (d *Dollar) Equals(o interface{}) bool {
	if value, ok := o.(*Dollar); ok {
		return *d == *value
	}

	return false
}
