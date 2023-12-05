package factory_method

import "testing"

func TestFactoryMethod(t *testing.T) {
	hwFactory := &HuaweiFactory{}
	hwPhone := hwFactory.CreatePhone()

	appleFactory := &AppleFactory{}
	iPhone := appleFactory.CreatePhone()

	hwPhone.Call()
	iPhone.Call()
}
