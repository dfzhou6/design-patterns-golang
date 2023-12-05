package state

import "fmt"

type Order struct {
	id    string
	state OrderState
}

func NewOrder(id string) *Order {
	return &Order{
		id:    id,
		state: &OrderCreateState{},
	}
}

func (impl *Order) SetState(state OrderState) {
	impl.state = state
}

func (impl *Order) Process() {
	impl.state.Handle(impl)
}

type OrderState interface {
	Handle(order *Order)
}

type OrderCreateState struct {
}

func (impl *OrderCreateState) Handle(order *Order) {
	fmt.Printf("order create state, order %+v\n", *order)
	order.SetState(&OrderPayingState{})
}

type OrderPayingState struct {
}

func (impl *OrderPayingState) Handle(order *Order) {
	fmt.Printf("order paying state, order %+v\n", *order)
	order.SetState(&OrderDoneState{})
}

type OrderDoneState struct {
}

func (impl *OrderDoneState) Handle(order *Order) {
	fmt.Printf("order done state, order %+v\n", *order)
}
