package decorator

import "fmt"

type IFavor interface {
	ShowFavor()
}

type Coffee struct {
	IFavor IFavor
}

func (c *Coffee) ShowFavor() {
	fmt.Println("this is coffee")
	if c.IFavor != nil {
		c.IFavor.ShowFavor()
	}
}

type Sugar struct {
	IFavor IFavor
}

func (s *Sugar) ShowFavor() {
	fmt.Println("add sugar")
	if s.IFavor != nil {
		s.IFavor.ShowFavor()
	}
}

type Milk struct {
	IFavor IFavor
}

func (i *Milk) ShowFavor() {
	fmt.Println("add milk")
	if i.IFavor != nil {
		i.IFavor.ShowFavor()
	}
}
