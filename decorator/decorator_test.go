package decorator

import "testing"

func TestNewCoffee(t *testing.T) {
	c := &Coffee{IFavor: &Sugar{IFavor: &Milk{}}}
	c.ShowFavor()
}
