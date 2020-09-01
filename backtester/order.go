package backtest

import (
	"github.com/thrasher-corp/gocryptotrader/currency"
	"github.com/thrasher-corp/gocryptotrader/exchanges/order"
)

func (o *Order) IsOrder() bool {
	return true
}

func (o *Order) SetDirection(s order.Side) {
	o.Direction = s
}

func (o *Order) GetDirection() order.Side {
	return o.Direction
}

func (o *Order) SetAmount(i float64) {
	o.Amount = i
}

func (o *Order) GetAmount() float64 {
	return o.Amount
}

func (o *Order) Pair() currency.Pair {
	return o.CurrencyPair
}

func (o *Order) Cancel() {
	o.Status = order.PendingCancel
}

func (o *Order) GetStatus() order.Status {
	return o.Status
}

func (o *Order) SetID(id int) {
	panic("implement me")
}

func (o *Order) ID() int {
	panic("implement me")
}

func (o *Order) GetLimit() float64 {
	panic("implement me")
}
