package money

// Money is a struct to express money.
type Money struct {
	amount   int
	currency string
}

// NewMoney creates new Money instance.
func NewMoney(amount int, currency string) *Money {
	return &Money{amount: amount, currency: currency}
}

// Times multiplies number to Money.
func (m *Money) Times(multiplier int) *Money {
	return NewMoney(m.amount*multiplier, m.currency)
}

// Plus pluses two Moneys.
func (m *Money) Plus(addend *Money) *Sum {
	return NewSum(m, addend)
}

// Reduce reduces Money.
func (m *Money) Reduce(to string) *Money {
	return m
}

// Equals compares two Moneys.
func (m *Money) Equals(m2 *Money) bool {
	return m.amount == m2.amount && m.currency == m2.currency
}

// Currency returns currency.
func (m *Money) Currency() string {
	return m.currency
}

// NewDollar is a constructor of Dollar.
func NewDollar(amount int) *Money {
	return &Money{amount: amount, currency: "USD"}
}

// NewFranc is a constructor of Franc.
func NewFranc(amount int) *Money {
	return &Money{amount: amount, currency: "CHF"}
}
