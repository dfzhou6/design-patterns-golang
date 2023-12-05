package state

import "testing"

func TestState(t *testing.T) {
	order := NewOrder("id-123456")
	order.Process()
	order.Process()
	order.Process()
	order.Process()
}
