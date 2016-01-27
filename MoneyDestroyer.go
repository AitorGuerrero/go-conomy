package goConomy

type MoneyDestroyer struct {
}

func (this MoneyDestroyer) ConsumeMoney(from MoneyGiver, m Money) (err error) {
	err, t := from.GiveTransaction(m)
	if nil != err {
		return
	}
	this.ReceiveTransaction(t)

	return
}

func (MoneyDestroyer) ReceiveTransaction(*transaction) {

}