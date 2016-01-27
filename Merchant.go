package goConomy

type MoneyReceiver interface {
	ReceiveTransaction(*transaction)
}

type MoneyGiver interface {
	GiveTransaction(Money) (error, *transaction)
}

type Merchant struct {
	wallet wallet
}

func (this *Merchant) GiveMoneyTo(amount Money, receiver MoneyReceiver) (err error) {
	err, t := this.GiveTransaction(amount)
	if (nil != err) {
		return;
	}
	receiver.ReceiveTransaction(t)

	return
}

func (this Merchant) HasEnoughMoney(m Money) bool {
	return this.wallet.totalAmount() >= m
}

func (this *Merchant) GiveTransaction(amount Money) (err error, t *transaction) {
	err, t = this.wallet.generateTransaction(amount)
	return
}

func (this *Merchant) ReceiveTransaction(t *transaction) {
	this.wallet.receiveTransaction(t)
}