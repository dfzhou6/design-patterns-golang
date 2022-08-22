package tpl

import (
	"errors"
	"fmt"
)

type ICall interface {
	call(number string) error
}

type Phone struct {
	ICall
}

func (*Phone) checkPhoneNumber(number string) error {
	if len(number) == 0 {
		return errors.New("phone number is empty")
	}
	fmt.Println("check phone number success")
	return nil
}

func (p *Phone) CallNumber(number string) error {
	var err error
	if err = p.checkPhoneNumber(number); err != nil {
		return err
	}
	return p.call(number)
}

type Huawei struct {
	*Phone
}

func (h *Huawei) call(number string) error {
	fmt.Println("huawei call number", number)
	return nil
}

func NewHuawei() *Huawei {
	huawei := &Huawei{}
	huawei.Phone = &Phone{
		ICall: huawei,
	}
	return huawei
}

type Samsung struct {
	*Phone
}

func (s *Samsung) call(number string) error {
	fmt.Println("samsung call number", number)
	return nil
}

func NewSamsung() *Samsung {
	samsung := &Samsung{}
	samsung.Phone = &Phone{
		ICall: samsung,
	}
	return samsung
}
