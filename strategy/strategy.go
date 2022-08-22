package strategy

import "fmt"

type IPay interface {
	Pay(limit string) error
}

var paymentMethod = map[string]IPay{
	"weixin": &WeiXinPay{},
	"ali":    &AliPay{},
}

type WeiXinPay struct {
}

func (*WeiXinPay) Pay(limit string) error {
	fmt.Println("WeiXin pay", limit)
	return nil
}

type AliPay struct {
}

func (*AliPay) Pay(limit string) error {
	fmt.Println("Ali pay", limit)
	return nil
}

func NewPaymentMethod(method string) IPay {
	return paymentMethod[method]
}
