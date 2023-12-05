package strategy

import "testing"

func TestStrategy(t *testing.T) {
	order := NewOrder(120)
	order.SetPayStrategy(&AliPay{})
	order.ProcessPayment()
	order.SetPayStrategy(&WeixinPay{})
	order.ProcessPayment()
}
