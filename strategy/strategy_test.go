package strategy

import "testing"

const limit = "200000000"

func TestAliPay_Pay(t *testing.T) {
	ali := NewPaymentMethod("ali")
	if err := ali.Pay(limit); err != nil {
		t.Errorf("ali Pay failed, err: %v", err)
	}
}

func TestWeiXinPay_Pay(t *testing.T) {
	wx := NewPaymentMethod("weixin")
	if err := wx.Pay(limit); err != nil {
		t.Errorf("wx Pay failed, err: %v", err)
	}
}
