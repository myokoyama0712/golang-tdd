package money

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	d := new(Dollar)
	d.amount = amount
	return d
}

func (d *Dollar) times(multiplier int) *Dollar {
	return NewDollar(d.amount * multiplier)
}

func (d *Dollar) equals(comp_d *Dollar) bool {
	return d.amount == comp_d.amount
}
