package goConomy

type MoneyGenerator struct {}

func (this MoneyGenerator) GenerateMoney(receiver MoneyReceiver, amount Money) {
	_, t := this.GiveTransaction(amount)
	receiver.ReceiveTransaction(t)
}

func (MoneyGenerator) GiveTransaction(amount Money) (error, *transaction) {
	return nil, &transaction{amount}
}