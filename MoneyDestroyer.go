package goConomy

type MoneyDestroyer struct {
}

func (this MoneyDestroyer) ConsumeMoney(from MoneyGiver, m Money) (err error) {
	err, t := from.giveTransaction(m)
	if nil != err {
		return
	}
	this.receiveTransaction(t)

	return
}

func (MoneyDestroyer) receiveTransaction(*transaction) {

}