package factory_method

import "fmt"

type Phone interface {
	Call()
}

type BasePhone struct {
	modelName string
}

func (impl *BasePhone) Call() {
	panic("pls implement me")
}

type HuaweiPhone struct {
	*BasePhone
}

func (impl *HuaweiPhone) Call() {
	fmt.Printf("this is huawei %v\n", impl.modelName)
}

type IPhone struct {
	*BasePhone
}

func (impl *IPhone) Call() {
	fmt.Printf("this is iPhone %v\n", impl.modelName)
}

type PhoneFactory interface {
	CreatePhone() Phone
}

type HuaweiFactory struct {
}

func (impl *HuaweiFactory) CreatePhone() Phone {
	return &HuaweiPhone{&BasePhone{"mate60 pro"}}
}

type AppleFactory struct {
}

func (impl *AppleFactory) CreatePhone() Phone {
	return &IPhone{&BasePhone{"15 pro max"}}
}
