package observer

import "fmt"

type Observer interface {
	GetID() string
	Update(data string)
}

type Customer struct {
	Name string
}

func NewCustomer(name string) *Customer {
	return &Customer{
		Name: name,
	}
}

func (impl *Customer) GetID() string {
	return impl.Name
}

func (impl *Customer) Update(data string) {
	fmt.Printf("customer %v receive data: %v\n", impl.Name, data)
}

type Subject interface {
	Register(observer Observer)
	UnRegister(observer Observer)
	NotifyAll()
}

type Product struct {
	Name        string
	Price       int
	customerMap map[string]Observer
}

func NewProduct(name string, price int) *Product {
	return &Product{
		Name:        name,
		Price:       price,
		customerMap: make(map[string]Observer, 0),
	}
}

func (impl *Product) Register(observer Observer) {
	impl.customerMap[observer.GetID()] = observer
}

func (impl *Product) UnRegister(observer Observer) {
	delete(impl.customerMap, observer.GetID())
}

func (impl *Product) NotifyAll() {
	for _, item := range impl.customerMap {
		item.Update(fmt.Sprintf("product name is %v, price is %v", impl.Name, impl.Price))
	}
}
