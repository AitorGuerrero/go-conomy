package goConomy

import t "testing"

func TestMultiply(t *t.T) {
	first := Money(5)
	result := first.Multiply(2)
	if result != 10 {
		t.Error("Should be 10")
	}
}

func TestAdd(t *t.T) {
	first := Money(5)
	result := first.Add(Money(2))
	if result != 7 {
		t.Error("Should be 7")
	}
}

func TestSubWhenUnder0(t *t.T) {
	first := Money(5)
	err, _ := first.Sub(Money(7))
	if err == nil {
		t.Error("Should return error")
	}
}

func TestSub(t *t.T) {
	first := Money(5)
	err, result := first.Sub(Money(2))
	if err != nil {
		t.Error(err)
	}
	if result != 3 {
		t.Error("The result should be 3", result)
	}
}