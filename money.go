package goConomy

import "errors"

type Money uint

const (
	Unit = Money(1)
)

func (this Money) Multiply(multiplier float32) Money {
	return Money(float32(this) * multiplier)
}

func (this Money) Add(amount Money) Money {
	return Money(this + amount)
}

func (this Money) Sub(amount Money) (err error, m Money) {
	if this < amount {
		return errors.New("Money can not  be negative"), Money(0)
	}
	return nil, this - amount
}