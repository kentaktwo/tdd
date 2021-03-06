package money

// Dollar is a struct to express dollar.
type Dollar struct {
	Money
}

// NewDollar is a constructor of Dollar.
func NewDollar(amount int) *Dollar {
	dollar := &Dollar{Money{amount: amount}}
	return dollar
}

// Times multiplies number to Dollar.
func (d *Dollar) Times(multiplier int) *Dollar {
	return NewDollar(d.amount * multiplier)
}

// Equals compares two Dollars.
func (d *Dollar) Equals(d2 *Dollar) bool {
	return d.Money.Equals(&d2.Money)
}
