package observer

import "testing"

func TestObserver(t *testing.T) {
	product := NewProduct("basketball", 300)
	alice := NewCustomer("alice")
	felix := NewCustomer("felix")

	product.Register(alice)
	product.Register(felix)
	product.NotifyAll()

	product.UnRegister(alice)
	product.NotifyAll()

	product.UnRegister(felix)
	product.NotifyAll()
}
