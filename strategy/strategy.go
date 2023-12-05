package strategy

import "fmt"

type PayStrategy interface {
	Pay(amount int)
}

type AliPay struct {
}

func (impl *AliPay) Pay(amount int) {
	fmt.Printf("use ali pay, amount %v\n", amount)
}

type WeixinPay struct {
}

func (impl *WeixinPay) Pay(amount int) {
	fmt.Printf("use weixin pay, amount %v\n", amount)
}

type Order struct {
	amount   int
	strategy PayStrategy
}

func NewOrder(amount int) *Order {
	return &Order{
		amount: amount,
	}
}

func (impl *Order) SetPayStrategy(strategy PayStrategy) {
	impl.strategy = strategy
}

func (impl *Order) ProcessPayment() {
	if impl.strategy == nil {
		panic("pls set pay strategy")
	}
	impl.strategy.Pay(impl.amount)
}
