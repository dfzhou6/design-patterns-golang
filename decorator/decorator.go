package decorator

type Coffee interface {
	GetCost() int
	GetDesc() string
}

type SimpleCoffee struct {
}

func (impl *SimpleCoffee) GetCost() int {
	return 10
}

func (impl *SimpleCoffee) GetDesc() string {
	return "coffee"
}

type SugarCoffee struct {
	coffee Coffee
}

func (impl *SugarCoffee) GetCost() int {
	return impl.coffee.GetCost() + 2
}

func (impl *SugarCoffee) GetDesc() string {
	return impl.coffee.GetDesc() + ", sugar"
}

type MilkCoffee struct {
	coffee Coffee
}

func (impl *MilkCoffee) GetCost() int {
	return impl.coffee.GetCost() + 3
}

func (impl *MilkCoffee) GetDesc() string {
	return impl.coffee.GetDesc() + ", milk"
}
