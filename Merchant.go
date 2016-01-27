package goConomy

type MoneyReceiver interface {
	receiveTransaction(*transaction)
}

type MoneyGiver interface {
	giveTransaction(Money) (error, *transaction)
}

type Merchant struct {
	wallet wallet
}

func (this *Merchant) GiveMoneyTo(amount Money, receiver MoneyReceiver) (err error) {
	err, t := this.giveTransaction(amount)
	if (nil != err) {
		return;
	}
	receiver.receiveTransaction(t)

	return
}

func (this Merchant) HasEnoughMoney(m Money) bool {
	return this.wallet.totalAmount() >= m
}

func (this *Merchant) giveTransaction(amount Money) (err error, t *transaction) {
	err, t = this.wallet.generateTransaction(amount)
	return
}

func (this *Merchant) receiveTransaction(t *transaction) {
	this.wallet.receiveTransaction(t)
}