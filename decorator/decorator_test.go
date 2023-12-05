package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	simpleCoffee := &SimpleCoffee{}
	sugarCoffee := &SugarCoffee{coffee: simpleCoffee}
	milkCoffee := &MilkCoffee{coffee: sugarCoffee}
	fmt.Printf("coffee cost: %v, desc: %v\n", milkCoffee.GetCost(), milkCoffee.GetDesc())
}
