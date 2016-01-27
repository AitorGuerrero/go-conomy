package goConomy

type MoneyGenerator struct {}

func (this MoneyGenerator) GenerateMoney(receiver MoneyReceiver, amount Money) {
	_, t := this.giveTransaction(amount)
	receiver.receiveTransaction(t)
}

func (MoneyGenerator) giveTransaction(amount Money) (error, *transaction) {
	return nil, &transaction{amount}
}